package metrics_collector

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var ClusterMetrics Metrics

func GetFlag() *string {
	configdir := flag.String("configdir", "", "absolute path to kubeconfig dir")
	flag.Parse()
	return configdir
}

func GetClusters() {
	//tmp := GetFlag()
	// kubeConfigs, err := ioutil.ReadDir(*GetFlag())
	//fmt.Printf("%s", *tmp)
	configPwd := GetFlag()
	kubeConfigs, err := ioutil.ReadDir(*configPwd)
	if err != nil {
		panic(err.Error())
	}
	ClusterMetrics = nil
	for _, kubeConfig := range kubeConfigs {
		if kubeConfig.IsDir() == false {
			config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(*configPwd, kubeConfig.Name()))
			if err != nil {
				ErrKubeconfig(&ClusterMetrics, kubeConfig.Name(), err)
				continue
			}
			cluster, err := kubernetes.NewForConfig(config)
			if err != nil {
				ErrKubeconfig(&ClusterMetrics, kubeConfig.Name(), err)
				continue
			} else {
				ClusterMetrics = append(ClusterMetrics, Metric{Cluster: cluster, Prov: kubeConfig.Name()})
			}
		}
	}
}
func ErrKubeconfig(clusterMetrics *Metrics, prov string, err error) {
	*clusterMetrics = append(*clusterMetrics, Metric{Cluster: nil, Prov: prov})
	panic(err.Error())
}
func CollectMetrics() {
	for i, _ := range ClusterMetrics {
		//Init node info
		ClusterMetrics[i].Analyses.NodeReadyNum = 0
		ClusterMetrics[i].Nodes = nil
		//Collect node info
		Nodes, err := ClusterMetrics[i].Cluster.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, node := range Nodes.Items {
			if node.Status.Conditions[3].Status == "True" {

				ClusterMetrics[i].Analyses.NodeReadyNum++
			}
			ClusterMetrics[i].Nodes = append(ClusterMetrics[i].Nodes,
				Node{Name: node.ObjectMeta.Name,
					Status: node.Status.Conditions[3],
					Host:   node.Status.Addresses[0].Address,
				})
			// fmt.Printf("%s", node.ObjectMeta.Name)
		}
		ClusterMetrics[i].Analyses.NodeNum = len(ClusterMetrics[i].Nodes)

		//Init pod info
		ClusterMetrics[i].Pods = nil
		ClusterMetrics[i].Analyses.PodRunningNum = 0
		//Collect pod info
		Pods, err := ClusterMetrics[i].Cluster.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error)
		}
		for _, pod := range Pods.Items {
			if pod.Status.Phase == "Running" {
				ClusterMetrics[i].Analyses.PodRunningNum++
			}
			ClusterMetrics[i].Pods = append(ClusterMetrics[i].Pods,
				Pod{
					Name:       pod.Name,
					NameSpaces: pod.Namespace,
					Status:     pod.Status.Phase,
					Endpoint:   pod.Status.PodIP,
					Host:       pod.Status.HostIP,
				})
		}
		ClusterMetrics[i].Analyses.PodNum = len(ClusterMetrics[i].Pods)

		//Init service info
		ClusterMetrics[i].Services = nil
		//Collect service info
		Services, err := ClusterMetrics[i].Cluster.CoreV1().Services("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, service := range Services.Items {
			ClusterMetrics[i].Services = append(ClusterMetrics[i].Services, Service{
				Name:      service.Name,
				Namespace: service.Namespace,
			})
		}
		ClusterMetrics[i].Analyses.ServiceNum = len(ClusterMetrics[i].Services)
		//Collect prov info
		ClusterMetrics[i].Analyses.Prov = ClusterMetrics[i].Prov
	}
}

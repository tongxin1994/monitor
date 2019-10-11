package restful_api

import (
	"encoding/json"
	"fmt"
	"monitor_api/metrics_collector"
	"net/http"

	"github.com/gorilla/mux"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
func ExportNodeInfo(w http.ResponseWriter, r *http.Request) {
	k8sClusters := metrics_collector.ClusterMetrics
	for _, k8sCluster := range k8sClusters {
		Nodes, err := k8sCluster.Cluster.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, node := range Nodes.Items {
			fmt.Printf("%s", node.ObjectMeta.Name)
			if err := json.NewEncoder(w).Encode(node.ObjectMeta.Name); err != nil {
				panic(err)
			}
		}
	}
}
func ExportPodInfo(w http.ResponseWriter, r *http.Request) {
	k8sClusters := metrics_collector.ClusterMetrics
	if err := json.NewEncoder(w).Encode(k8sClusters[0].Pods); err != nil {
		panic(err)
	}
}

func ExportClusterInfo(w http.ResponseWriter, r *http.Request) {
	metrics_collector.CollectMetrics()
	for i, _ := range metrics_collector.ClusterMetrics {
		if err := json.NewEncoder(w).Encode(metrics_collector.ClusterMetrics[i].Analyses); err != nil {
			panic(err)
		}
		tmp, _ := metrics_collector.ClusterMetrics[i].Cluster.CoreV1().Nodes().List(metav1.ListOptions{})
		for _, i := range tmp.Items {
			if err := json.NewEncoder(w).Encode(i.Status); err != nil {
				panic(err)
			}
		}
	}
}

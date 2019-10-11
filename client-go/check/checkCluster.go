package check

// import (
// 	"client-go/kubeconfig"
// 	"client-go/log"
// 	"os"

// 	"k8s.io/client-go/kubernetes"
// )

// //CheckCluster check pod and node status of given configfile, generate check logs
// func CheckCluster(clientset *kubernetes.Clientset, file *os.File, kubeConfig string) error {
// 	clientset, err := kubeconfig.MakeClientSet(file, kubeConfig.Name())
// 	if err != nil {
// 		continue
// 	}
// 	if check.GetNode(clientset, file, kubeConfig.Name()) != nil {
// 		continue
// 	}
// 	if check.GetPod(clientset, file, kubeConfig.Name()) != nil {
// 		continue
// 	}
// 	if check.NodeReady != len(check.Nodes.Items) || check.PodRunning != len(check.Pods.Items) {
// 		log.WriteLog(file, "                  Details                     \n")
// 	}
// 	check.NodeDetails(clientset, file)
// 	check.PodDetails(clientset, file)
// }

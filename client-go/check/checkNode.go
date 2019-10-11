package check

import (
	"client-go/log"
	"fmt"
	"os"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//NodeReady count ready node num
var NodeReady int

//Nodes is node list
var Nodes *v1.NodeList

//GetNode list all nodes stations in a cluster
func GetNode(clientset *kubernetes.Clientset, file *os.File, kubeConfig string) error {
	var err error
	Nodes, err = clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		// panic(err.Error())
		log.WriteLog(file, fmt.Sprintf("*********************************%10s*************************************\n", kubeConfig))
		log.WriteLog(file, "Check kubeconfig file!\n")
		return err
	}

	log.WriteLog(file, fmt.Sprintf("*********************************%10s*************************************\n", kubeConfig))
	log.WriteLog(file, "                  Summary                     \n")
	log.WriteLog(file, "┌────────────┬──────────────┬────────────────┐\n")
	log.WriteLog(file, "│            │     Total    │ Running(Active)│\n")
	log.WriteLog(file, "├────────────┼──────────────┼────────────────┤\n")

	NodeReady = 0

	for _, node := range Nodes.Items {
		if node.Status.Conditions[3].Status == "True" {
			NodeReady++
		}
	}
	log.WriteLog(file, fmt.Sprintf("│   Nodes    │     %4d     │   %4d/%-4d    │\n", len(Nodes.Items), NodeReady, len(Nodes.Items)))
	log.WriteLog(file, "├────────────┼──────────────┼────────────────┤\n")
	return err
}

//NodeDetails shows node not ready details
func NodeDetails(clientset *kubernetes.Clientset, file *os.File) {
	if NodeReady != len(Nodes.Items) {
		log.WriteLog(file, "Node not Ready details:\n")
		for _, node := range Nodes.Items {
			if node.Status.Conditions[3].Status != "True" {
				log.WriteLog(file, fmt.Sprintf("Name: %s, ip: %s, Events: %s \n", node.ObjectMeta.Name, node.Status.Addresses[0].Address, node.Status.Conditions[4].Message))
			}
		}
	}
}

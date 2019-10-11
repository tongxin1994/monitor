package check

import (
	"client-go/log"
	"fmt"
	"os"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//PodRunning counts the number of running pods
var PodRunning int

//Pods is type of pods
var Pods *v1.PodList

//GetPod lists pod stations in the cluster
func GetPod(clientset *kubernetes.Clientset, file *os.File, kubeConfig string) error {
	var err error
	Pods, err = clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	PodRunning = 0
	for _, pod := range Pods.Items {
		if pod.Status.Phase == "Running" {
			PodRunning++
		}
	}
	log.WriteLog(file, fmt.Sprintf("│   Pods     │     %4d     │   %4d/%-4d    │\n", len(Pods.Items), PodRunning, len(Pods.Items)))
	log.WriteLog(file, "└────────────┴──────────────┴────────────────┘\n")
	return err
}

//PodDetails shows not running pod details
func PodDetails(clientset *kubernetes.Clientset, file *os.File) {
	if PodRunning != len(Pods.Items) {
		log.WriteLog(file, "Pod not Running details: \n")
		for _, pod := range Pods.Items {
			if pod.Status.Phase != "Running" {
				log.WriteLog(file, fmt.Sprintf("Name: %s, Namespace: %s, node: %s, Status: %s \n", pod.ObjectMeta.Name, pod.ObjectMeta.Namespace, pod.Spec.NodeName, pod.Status.Phase))
			}
		}
	}
}

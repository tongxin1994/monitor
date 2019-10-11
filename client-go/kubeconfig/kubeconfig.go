package kubeconfig

import (
	"client-go/getflag"
	"client-go/log"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//MakeClientSet create clinetset with kubeconfig files
func MakeClientSet(file *os.File, kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(getflag.ConfigDir, kubeconfig))
	if err != nil {
		log.WriteLog(file, fmt.Sprintf("*********************************%10s*************************************\n", kubeconfig))
		log.WriteLog(file, "Check kubeconfig file!\n")
		return nil, err
	} else {
		// create the clientset
		return kubernetes.NewForConfig(config)
	}
}

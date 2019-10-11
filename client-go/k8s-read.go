/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"client-go/check"
	"client-go/getflag"
	"client-go/kubeconfig"
	"client-go/log"
	"io/ioutil"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	//get flag from command lines
	getflag.GetFlag()
	//create log file
	file, err := log.CreateLogFile()
	if err != nil {
		panic(err.Error())
	}

	configs, err := ioutil.ReadDir(getflag.ConfigDir)
	if err != nil {
		panic(err.Error())
	}

	for _, kubeConfig := range configs {
		// go func (kubeConfig os.FileInfo) {
		if kubeConfig.IsDir() {
			continue
		} else {
			clientset, err := kubeconfig.MakeClientSet(file, kubeConfig.Name())
			if err != nil {
				continue
			}
			if check.GetNode(clientset, file, kubeConfig.Name()) != nil {
				continue
			}
			if check.GetPod(clientset, file, kubeConfig.Name()) != nil {
				continue
			}
			if check.NodeReady != len(check.Nodes.Items) || check.PodRunning != len(check.Pods.Items) {
				log.WriteLog(file, "                  Details                     \n")
			}
			check.NodeDetails(clientset, file)
			check.PodDetails(clientset, file)
		}
	}
}

// }

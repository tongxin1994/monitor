package metrics_collector

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Metric struct {
	Cluster  *kubernetes.Clientset `json:"cluster"`
	Nodes    []Node                `json:"nodes"`
	Pods     []Pod                 `json:"pods"`
	Services []Service             `json:"services"`
	Deploys  []Deployment          `json:"deploys"`
	Analyses AnalyzedMetrics       `json:"cluster-info"`
	Prov     string                `json:"prov"`
}
type Metrics []Metric

type Node struct {
	Name   string           `json:"name"`
	Status v1.NodeCondition `json:"status"`
	Host   string           `json:"host"`
}
type Pod struct {
	Name       string      `json:"name"`
	NameSpaces string      `json:"namespaces"`
	Status     v1.PodPhase `json:"status"`
	Endpoint   string      `json:"endpoint"`
	Host       string      `json:"host"`
}
type Service struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	ClusterIP string `json:"clusterip"`
	NodePort  string `json:"nodeport"`
}
type Deployment struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicate int    `json:"repalicate"`
}
type AnalyzedMetrics struct {
	Prov          string `json:"prov"`
	NodeNum       int    `json:"node-sum"`
	NodeReadyNum  int    `json:"node-ready"`
	PodNum        int    `json:"pod-sum"`
	PodRunningNum int    `json:"pod-running"`
	ServiceNum    int    `json:"service-sum"`
}

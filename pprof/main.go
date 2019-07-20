package main

import (
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
)

//service cache obj
type K8sObj struct {
	Deployment  []v1beta1.Deployment  `json:"deployment"`
	Pod         *v1.PodList           `json:"pod"`
	DaemonSet   v1beta2.DaemonSet     `json:"daemon_set"`
	StatefulSet []v1beta1.StatefulSet `json:"stateful_set"`
}

func main() {

}

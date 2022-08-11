package pod

import (
	"github.com/go-kit/log"
	"k8s.io/client-go/kubernetes"
)

type Pod struct {
	clientset *kubernetes.Clientset
	PodName   string
	NameSpace string
	// v1pod     *apiv1.Pod

	logger log.Logger
}

const (
	ContainersReady string = "ContainersReady"
	PodInitialized  string = "Initialized"
	PodReady        string = "Ready"
	PodScheduled    string = "PodScheduled"
)

const (
	ConditionTrue    string = "True"
	ConditionFalse   string = "False"
	ConditionUnknown string = "Unknown"
)

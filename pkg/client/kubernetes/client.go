package kubernetes

import (
	"io/ioutil"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewRestClient() (*kubernetes.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset, nil
}

func NewClient(kubeConfig string) (*kubernetes.Clientset, error) {
	data, err := ioutil.ReadFile(kubeConfig)
	if err != nil {
		return nil, err
	}
	config, err := clientcmd.NewClientConfigFromBytes(data)
	if err != nil {
		return nil, err
	}
	restConfig, err := config.ClientConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

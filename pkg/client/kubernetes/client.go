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

	// for {
	// 	// get pods in all the namespaces by omitting namespace
	// 	// Or specify namespace to get pods in particular namespace
	// 	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// 	// Examples for error handling:
	// 	// - Use helper functions e.g. errors.IsNotFound()
	// 	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// 	_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
	// 	if errors.IsNotFound(err) {
	// 		fmt.Printf("Pod example-xxxxx not found in default namespace\n")
	// 	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 		fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
	// 	} else if err != nil {
	// 		panic(err.Error())
	// 	} else {
	// 		fmt.Printf("Found example-xxxxx pod in default namespace\n")
	// 	}

	// 	time.Sleep(10 * time.Second)
	// }
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

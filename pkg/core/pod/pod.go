package pod

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreatePod(clientset *kubernetes.Clientset, env_gitsource string, env_callback string) {
	// 创建一个GitOps Pod
	namespace := "default"
	podsClient := clientset.CoreV1().Pods(namespace)
	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-gitops",
			Labels: map[string]string{
				"app": "gitops",
			},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:  "web",
					Image: "nginx:1.12",
					Ports: []apiv1.ContainerPort{
						{
							Name:          "http",
							Protocol:      apiv1.ProtocolTCP,
							ContainerPort: 80,
						},
					},
				},
			},
			RestartPolicy: apiv1.RestartPolicyNever,
		},
	}
	// Create Pod
	fmt.Println("Creating gitops pod...")
	result, err := podsClient.Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created gitops pod %q.\n", result.GetObjectMeta().GetName())

	// GitOps Controller，Pod运行完成删除Pod

	//List && Watch GitOps Pod

	// // 创建一个Deployment
	// deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	// deployment := &appsv1.Deployment{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: "demo-gitops",
	// 	},
	// 	Spec: appsv1.DeploymentSpec{
	// 		Replicas: int32Ptr(1),
	// 		Selector: &metav1.LabelSelector{
	// 			MatchLabels: map[string]string{
	// 				"app": "gitops",
	// 			},
	// 		},
	// 		Template: apiv1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Labels: map[string]string{
	// 					"app": "gitops",
	// 				},
	// 			},
	// 			Spec: apiv1.PodSpec{
	// 				Containers: []apiv1.Container{
	// 					{
	// 						Name:  "web",
	// 						Image: "nginx:1.12",
	// 						Ports: []apiv1.ContainerPort{
	// 							{
	// 								Name:          "http",
	// 								Protocol:      apiv1.ProtocolTCP,
	// 								ContainerPort: 80,
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// // Create Deployment
	// fmt.Println("Creating gitops deployment...")
	// result2, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Created gitops deployment %q.\n", result2.GetObjectMeta().GetName())

}

// func int32Ptr(i int32) *int32 { return &i }

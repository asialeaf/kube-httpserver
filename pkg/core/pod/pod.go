package pod

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GitOps Controller，Pod运行完成删除Pod

func CreatePod(clientset *kubernetes.Clientset, env_gitsource, env_callback, env_gitpath string) (*apiv1.Pod, error) {

	// 创建一个GitOps Pod
	namespace := "default"
	gitOpsPodName := "demo-gitops"
	podsClient := clientset.CoreV1().Pods(namespace)
	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: gitOpsPodName,
			Labels: map[string]string{
				"app": "gitops",
			},
		},
		Spec: apiv1.PodSpec{
			HostAliases: []apiv1.HostAlias{
				{
					IP:        "192.168.50.225",
					Hostnames: []string{"cluster-endpoint"},
				},
			},
			Containers: []apiv1.Container{
				{
					Name:  "gitops",
					Image: "kubectl:v1.0",
					Ports: []apiv1.ContainerPort{
						{
							Name:          "http",
							Protocol:      apiv1.ProtocolTCP,
							ContainerPort: 80,
						},
					},
					Command:         []string{"sh", "-c", "/gitops.sh"},
					ImagePullPolicy: apiv1.PullNever,
					Env: []apiv1.EnvVar{
						{
							Name:  "GITSOURCE",
							Value: env_gitsource,
						},
						{
							Name:  "GITPATH",
							Value: env_gitpath,
						},
						{
							Name:  "CALLBACK",
							Value: env_callback,
						},
					},
					Resources: apiv1.ResourceRequirements{},
					VolumeMounts: []apiv1.VolumeMount{
						{
							Name:      "kubeconfig",
							ReadOnly:  true,
							MountPath: "/root/.kube",
						},
					},
				},
			},
			RestartPolicy: apiv1.RestartPolicyNever,
			NodeSelector: map[string]string{
				"kubernetes.io/hostname": "k8s-master",
			},
			Tolerations: []apiv1.Toleration{
				{
					Effect:   apiv1.TaintEffectNoSchedule,
					Operator: apiv1.TolerationOpExists,
				},
			},
			Volumes: []apiv1.Volume{
				{
					Name: "kubeconfig",
					VolumeSource: apiv1.VolumeSource{
						HostPath: &apiv1.HostPathVolumeSource{
							Path: "/root/.kube",
						},
					},
				},
			},
		},
	}
	// Create Pod
	fmt.Println("Creating gitops pod...")
	result, err := podsClient.Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created gitops pod %q.\n", result.GetObjectMeta().GetName())

	return result, err
}

// List && Watch GitOps Pod

// Get Pod
func GetPod(clientset *kubernetes.Clientset, podName, namespace string) (*apiv1.Pod, error) {
	return clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
}

// 获取Pod状态
func GetPodStatus(pod *apiv1.Pod) string {
	// for _, cond := range pod.Status.Conditions {
	// 	if string(cond.Type) == ContainersReady {
	// 		if string(cond.Status) != ConditionTrue {
	// 			return "Unavailable"
	// 		}
	// 	} else if string(cond.Type) == PodInitialized && string(cond.Status) != ConditionTrue {
	// 		return "Initializing"
	// 	} else if string(cond.Type) == PodReady {
	// 		if string(cond.Status) != ConditionTrue {
	// 			return "Unavailable"
	// 		}
	// 		for _, containerState := range pod.Status.ContainerStatuses {
	// 			if !containerState.Ready {
	// 				return "Unavailable"
	// 			}
	// 		}
	// 	} else if string(cond.Type) == PodScheduled && string(cond.Status) != ConditionTrue {
	// 		return "Scheduling"
	// 	}
	// }
	return string(pod.Status.Phase)
}

//Delete Pod
func DeletePod(clientset *kubernetes.Clientset, podName, namespace string) {
	fmt.Println("Deleting Pod...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted Pod.")
}

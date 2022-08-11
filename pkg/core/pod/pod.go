package pod

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func New(logger log.Logger, clientset *kubernetes.Clientset, name string, namespace string) *Pod {
	if logger == nil {
		logger = log.NewNopLogger()
	}
	p := &Pod{
		logger:    logger,
		clientset: clientset,
		PodName:   name,
		NameSpace: namespace,
	}
	return p
}

func (p *Pod) CreatePod(env_gitsource, env_gitpath, env_callback string) (*apiv1.Pod, error) {
	logger := log.With(p.logger, "pod", p.PodName)

	// 创建一个GitOps Pod
	podsClient := p.clientset.CoreV1().Pods(p.NameSpace)
	v1pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: p.PodName,
			Labels: map[string]string{
				"app": p.PodName,
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
					Name:  p.PodName,
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
	level.Info(logger).Log("msg", "Creating pod...")
	result, err := podsClient.Create(context.TODO(), v1pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	level.Info(logger).Log("msg", fmt.Sprintf("Created pod %q", result.GetObjectMeta().GetName()))
	// p.v1pod = result
	return result, err
}

// Get Pod
func (p *Pod) GetPod() (*apiv1.Pod, error) {
	return p.clientset.CoreV1().Pods(p.NameSpace).Get(context.TODO(), p.PodName, metav1.GetOptions{})
}

// Get Pod Status
func (p *Pod) GetPodStatus(pod *apiv1.Pod) string {
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

// Delete Pod
func (p *Pod) DeletePod() {
	logger := log.With(p.logger, "pod", p.PodName)
	level.Info(logger).Log("msg", "Deleting pod...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := p.clientset.CoreV1().Pods(p.NameSpace).Delete(context.TODO(), p.PodName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	level.Info(logger).Log("msg", "Deleted Pod.")
}

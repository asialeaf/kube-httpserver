package controller

import (
	"fmt"
	"time"

	"git.harmonycloud.cn/yeyazhou/kube-httpserver/pkg/client/kubernetes"
	"git.harmonycloud.cn/yeyazhou/kube-httpserver/pkg/core/pod"
)

// GitOps Controller，Pod运行完成删除Pod
func Handler(gitsource, gitpath, callback string) {

	// client, _ := kubernetes.NewRestClient()
	client, _ := kubernetes.NewClient("/root/.kube/config")
	gitopspod, err := pod.CreatePod(client, gitsource, gitpath, callback)
	if err != nil {
		panic(err)
	}
	podName := gitopspod.GetObjectMeta().GetName()
	podNamaSpace := gitopspod.GetObjectMeta().GetNamespace()

	for {
		time.Sleep(5 * time.Second)
		gitOpsPod, _ := pod.GetPod(client, podName, podNamaSpace)
		podStatus := pod.GetPodStatus(gitOpsPod)

		fmt.Printf("%s pod status: %s\n", podName, podStatus)

		if podStatus == "Failed" {
			fmt.Printf("%s run %s\n", podName, podStatus)
			pod.DeletePod(client, podName, podNamaSpace)
			break
		} else if podStatus == "Succeeded" {
			// fmt.Printf("%s run %s\n", podName, podStatus)
			pod.DeletePod(client, podName, podNamaSpace)
			break
		}

	}

	// if podStatus == "Succeeded" {
	// 	pod.DeletePod(client, "demo-gitops", "default")
	// }
}

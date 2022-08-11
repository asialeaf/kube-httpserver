package controller

import (
	"fmt"
	stdlog "log"
	"time"

	"git.harmonycloud.cn/gitops/kube-httpserver/pkg/client/kubernetes"
	"git.harmonycloud.cn/gitops/kube-httpserver/pkg/core/pod"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

// GitOps Controller，Pod运行完成删除Pod
func Handler(gitsource, gitpath, callback string, logger log.Logger) {

	// client, _ := kubernetes.NewRestClient()
	client, _ := kubernetes.NewClient("/root/.kube/config")

	p := pod.New(logger, client, "demo-gitops", "default")
	_, err := p.CreatePod(gitsource, gitpath, callback)
	if err != nil {
		stdlog.Fatalf("gitops pod creation failed, err: %s", err)
	}

	for {
		time.Sleep(5 * time.Second)
		gitOpsPod, _ := p.GetPod()
		podStatus := p.GetPodStatus(gitOpsPod)
		level.Info(logger).Log("msg", fmt.Sprintf("%s pod status: %s.", p.PodName, podStatus))
		if podStatus == "Failed" {
			level.Warn(logger).Log("msg", fmt.Sprintf("%s run %s, gitops pod is about to be deleted.", p.PodName, podStatus))
			p.DeletePod()
			break
		} else if podStatus == "Succeeded" {
			level.Info(logger).Log("msg", fmt.Sprintf("%s run %s, gitops pod is about to be deleted.", p.PodName, podStatus))
			p.DeletePod()
			break
		}
	}
}

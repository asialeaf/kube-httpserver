package server

import (
	"fmt"
	"net/http"
	"time"

	"git.harmonycloud.cn/yeyazhou/kube-httpserver/pkg/client/kubernetes"
	"git.harmonycloud.cn/yeyazhou/kube-httpserver/pkg/core/pod"
	"github.com/gin-gonic/gin"
)

func Demo() {
	router := gin.Default()

	router.POST("/deploy", func(c *gin.Context) {
		var json Data
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//参数校验
		// if json.GitSource != "manu" || json.CallBack != "123" {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"status": "参数错误"})
		// 	return
		// }

		//业务逻辑
		// client, _ := kubernetes.NewRestClient()
		client, _ := kubernetes.NewClient("/root/.kube/config")
		_, err := pod.CreatePod(client, json.GitSource, json.CallBack)
		if err != nil {
			panic(err)
		}

		time.Sleep(10 * time.Second)
		gitOpsPod, _ := pod.GetPod(client, "demo-gitops", "default")
		podStatus := pod.GetPodStatus(gitOpsPod)

		fmt.Println(podStatus)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	router.Run(":8080")
}

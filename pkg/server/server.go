package server

import (
	"net/http"

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
		client, _ := kubernetes.NewRestClient()
		pod.CreatePod(client, json.GitSource, json.CallBack)

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	router.Run(":8080")
}

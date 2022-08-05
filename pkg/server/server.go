package server

import (
	"net/http"

	"git.harmonycloud.cn/gitops/kube-httpserver/pkg/core/controller"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Demo() {
	router := gin.Default()

	router.POST("/gitops/deploy", func(c *gin.Context) {
		var json Data
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//参数校验
		if !govalidator.IsURL(json.GitSource) || !govalidator.IsURL(json.CallBack) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "参数错误"})
			return
		}

		//业务逻辑
		go controller.Handler(json.GitSource, json.GitPath, json.CallBack)
		c.JSON(http.StatusOK, gin.H{"status": "success"})

	})

	router.Run(":8080")
}

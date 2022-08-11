package server

import (
	"fmt"
	"net/http"

	"git.harmonycloud.cn/gitops/kube-httpserver/pkg/core/controller"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func Demo(logger log.Logger) {
	if logger == nil {
		logger = log.NewNopLogger()
	}
	level.Info(logger).Log("msg", "Starting kube-httpserver")
	gin.SetMode(gin.ReleaseMode)
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
		level.Info(logger).Log("msg", fmt.Sprintf("The requested cluster is %s", json.ClusterName))
		//业务逻辑
		go controller.Handler(json.GitSource, json.GitPath, json.CallBack, log.With(logger, "component", "controller"))
		c.JSON(http.StatusOK, gin.H{"status": "success"})

	})
	level.Info(logger).Log("msg", fmt.Sprintf("Kube-httpserver started. Listening and serving HTTP on %s", ListenAddress))
	router.Run(ListenAddress)
}

package main

import (
	"git.harmonycloud.cn/gitops/kube-httpserver/pkg/server"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/common/promlog"
)

func main() {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	level.Info(logger).Log("msg", "Starting kube-httpserver")
	webLogger := log.With(logger, "component", "web")
	server.Demo(webLogger)
}

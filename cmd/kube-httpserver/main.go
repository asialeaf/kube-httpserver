package main

import (
	"git.harmonycloud.cn/gitops/kube-httpserver/pkg/server"
	"github.com/go-kit/log"
	"github.com/prometheus/common/promlog"
)

func main() {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	webLogger := log.With(logger, "component", "web")
	server.Demo(webLogger)
}

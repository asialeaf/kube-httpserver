package main

import (
	"fmt"

	"git.harmonycloud.cn/gitops/kube-httpserver/pkg/server"
)

func main() {
	// fmt.Println("clinet-go 测试...")
	// kubeclient.ClientTest()
	fmt.Println("gitops httpserver ...")
	server.Demo()
}

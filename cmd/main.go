package main

import (
	"fmt"

	"git.harmonycloud.cn/yeyazhou/kube-httpserver/pkg/kubeclient"
)

func main() {
	fmt.Println("clinet-go 测试...")
	kubeclient.ClientTest()
}

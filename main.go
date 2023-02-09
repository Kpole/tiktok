// Code generated by hertz generator.

package main

import (
	"offer_tiktok/biz/dal"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	h := server.Default(
		server.WithHostPorts("0.0.0.0:18005"),
	)

	register(h)
	h.Spin()
}

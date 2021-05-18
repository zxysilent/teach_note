package main

import (
	"noteapi/conf"
	"noteapi/model"
	"noteapi/router"
	"os"
	"os/signal"
	"syscall"

	"github.com/zxysilent/logs"
)

// @Title noteapi’s Api文档
// @Version 1.0
// @Description 凭证传递方式包括 get、post、header、cookie
// @Description 数据传递方式包括 json、formData
// @Description /api/* 公共访问
// @Description /adm/* 必须传入token
func main() {
	logs.Info("app initializing")
	conf.Init()
	model.Init()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	logs.Info("app running")
	go router.RunApp()
	<-quit
	logs.Info("app quitted")
	logs.Flush()
}

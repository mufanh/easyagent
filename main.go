package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/routers"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

var (
	configPath string
	configName string
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/", "配置文件加载路径")
	flag.StringVar(&configName, "configName", "server", "配置文件名")
	flag.Parse()

	if err := global.SetupServerSetting(configPath, configName); err != nil {
		panic(errors.Wrap(err, "初始化应用配置失败"))
	}
	if err := global.SetupLogger(
		global.ServerLogConfig.LogFilepath,
		global.ServerLogConfig.LogFilename,
		global.ServerLogConfig.LogMaxSize,
		global.ServerLogConfig.LogMaxAge); err != nil {
		panic(errors.Wrap(err, "初始化应用日志失败"))
	}

}

// @title easyagent
// @version 1.0
// @description Agent控制管理器
// @termsOfService https://github.com/mufanh/easyagent
func main() {
	gin.SetMode(global.ServerConfig.RunMode)
	router := routers.NewServerRouter()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(int(global.ServerConfig.HttpPort)),
		Handler:        router,
		ReadTimeout:    global.ServerConfig.ReadTimeout,
		WriteTimeout:   global.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		global.Logger.Panicf("应用Http服务启动失败，详细错误信息:%+v\n", err)
	}
}

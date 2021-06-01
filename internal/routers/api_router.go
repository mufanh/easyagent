package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mufanh/easyagent/docs"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/middleware"
	"github.com/mufanh/easyagent/internal/routers/api"
	"github.com/mufanh/easyagent/internal/routers/ws"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewServerRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.ResponseLog())
	if global.ServerConfig.RunMode == "debug" {
		r.Use(middleware.Cors())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.ContextTimeout(global.ServerConfig.Timeout))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/websocket", func(c *gin.Context) {
		if err := ws.Connect(c); err != nil {
			global.Logger.Warnf("连接失败，失败详细原因:%+v", err)
			http.Error(c.Writer, err.Error(), http.StatusForbidden)
			return
		}
	})

	sessionRouter := api.SessionApiRouter{}
	r.GET("/api/sessions", sessionRouter.List)
	r.DELETE("/api/sessions/:token", sessionRouter.Close)

	shellRouter := api.ShellApiRouter{}
	r.POST("/api/shell/exec", shellRouter.Exec)
	r.POST("/api/shell/log", shellRouter.ShowLog)

	scriptRouter := api.ScriptApiRouter{}
	r.POST("/api/script/upload", scriptRouter.Upload)
	r.POST("/api/script/exec", scriptRouter.Exec)
	r.POST("/api/script/log", scriptRouter.ShowLog)
	r.POST("/api/script/groups", scriptRouter.ShowGroupDirs)
	r.POST("/api/script/files", scriptRouter.ShowScriptFiles)

	return r
}

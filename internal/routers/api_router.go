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

	sessionRouter := new(api.SessionApiRouter)
	r.GET("/api/sessions", sessionRouter.List)

	commandRouter := new(api.CommandApiRouter)
	r.POST("/api/command/exec", commandRouter.Exec)

	scriptRouter := new(api.ScriptApiRouter)
	r.POST("/api/script/upload", scriptRouter.Upload)
	r.POST("/api/script/show", scriptRouter.Show)
	r.POST("/api/script/delete", scriptRouter.Delete)
	r.POST("/api/script/group/delete", scriptRouter.DeleteGroupDir)
	r.POST("/api/script/exec", scriptRouter.Exec)
	r.POST("/api/script/groups", scriptRouter.ShowGroupDirs)
	r.POST("/api/script/files", scriptRouter.ShowScriptFiles)

	shellLogRouter := new(api.ShellLogApiRouter)
	r.POST("/api/shell/log", shellLogRouter.Show)

	return r
}

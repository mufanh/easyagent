package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mufanh/easyagent/docs"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/routers/ws"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/websocket", func(c *gin.Context) {
		if err := ws.Connect(c); err != nil {
			global.Logger.Warnf("连接失败，失败详细原因:%+v", err)
			http.Error(c.Writer, err.Error(), http.StatusForbidden)
			return
		}
	})
	return r
}

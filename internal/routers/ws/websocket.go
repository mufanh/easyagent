package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 5 * time.Second,
	ReadBufferSize:   2048,
	WriteBufferSize:  2048,
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// agent通过websocket连接server
func Connect(c *gin.Context) error {
	data := make(map[string]string)
	if err := c.BindHeader(data); err != nil {
		return err
	}

	agentInfo := model.ConvertMap2AgentInfo(data)
	if err := agentInfo.Check(); err != nil {
		return err
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}

	if err := global.ServerRepo.AddSession(agentInfo, conn); err != nil {
		return err
	}

	return nil
}

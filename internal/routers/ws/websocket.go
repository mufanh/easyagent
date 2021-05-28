package ws

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/issue9/jsonrpc"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"net/http"
	"strings"
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
	if len(c.Request.Header) > 0 {
		for k, v := range c.Request.Header {
			data[strings.ToUpper(k)] = v[0]
		}
	}
	agentInfo := model.ConvertMap2AgentInfo(data)
	if err := agentInfo.Check(); err != nil {
		return err
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}

	t := jsonrpc.NewWebsocketTransport(conn)
	serv := jsonrpc.NewServer()
	jConn := serv.NewConn(t, nil)

	if err := global.ServerRepo.AddSession(agentInfo, conn, jConn); err != nil {
		return err
	}
	defer func() {
		if err := global.ServerRepo.DeleteSession(agentInfo.Token); err != nil {
			global.Logger.Warnf("删除连接失败，详细错误原因:%+V", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := jConn.Serve(ctx); err != nil {
		return err
	}

	return nil
}

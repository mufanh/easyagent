package global

import (
	"github.com/gorilla/websocket"
	"github.com/issue9/jsonrpc"
)

var transport jsonrpc.Transport
var conn *websocket.Conn

func SetConn(c *websocket.Conn) {
	conn = c
	transport = jsonrpc.NewWebsocketTransport(conn)
}

func GetConn() *websocket.Conn {
	return conn
}

func GetTransport() jsonrpc.Transport {
	return transport
}

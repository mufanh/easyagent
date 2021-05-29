package routers

import (
	"github.com/issue9/jsonrpc"
	jsonrpc2 "github.com/mufanh/easyagent/internal/routers/jsonrpc"
)

func NewAgentJsonRpcRouter() *jsonrpc.Server {
	server := new(jsonrpc.Server)

	sessionRouter := new(jsonrpc2.SessionJsonRpcRouter)
	server.Register("session.close", sessionRouter.Close)

	shellRouter := new(jsonrpc2.ShellJsonRpcRouter)
	server.Register("shell.exec", shellRouter.Exec)
	server.Register("shell.log", shellRouter.ShowLog)

	scriptRouter := new(jsonrpc2.ScriptJsonRpcRouter)
	server.Register("script.upload", scriptRouter.Upload)
	server.Register("script.exec", scriptRouter.Exec)
	server.Register("script.log", scriptRouter.ShowLog)

	return server
}

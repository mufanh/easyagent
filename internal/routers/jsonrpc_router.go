package routers

import (
	"github.com/issue9/jsonrpc"
	jsonrpc2 "github.com/mufanh/easyagent/internal/routers/jsonrpc"
)

func NewAgentJsonRpcRouter() *jsonrpc.Server {
	server := new(jsonrpc.Server)

	shellRouter := new(jsonrpc2.CommandJsonRpcRouter)
	server.Register("shell.exec", shellRouter.Exec)

	scriptRouter := new(jsonrpc2.ScriptJsonRpcRouter)
	server.Register("script.upload", scriptRouter.Upload)
	server.Register("script.show", scriptRouter.Show)
	server.Register("script.delete", scriptRouter.Delete)
	server.Register("script.update", scriptRouter.Update)
	server.Register("script.deleteGroupDir", scriptRouter.DeleteGroupDir)
	server.Register("script.exec", scriptRouter.Exec)
	server.Register("script.showGroupDirs", scriptRouter.ShowGroupDirs)
	server.Register("script.showScriptFiles", scriptRouter.ShowScriptFiles)

	shellLogRouter := new(jsonrpc2.ShellLogJsonRpcRouter)
	server.Register("shell.showLog", shellLogRouter.Show)

	return server
}

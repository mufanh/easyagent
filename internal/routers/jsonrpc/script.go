package jsonrpc

import "github.com/mufanh/easyagent/internal/model"

type ScriptJsonRpcRouter struct {
}

func (s ScriptJsonRpcRouter) Upload(notify bool, request *model.ScriptUploadRequest, response *model.ScriptUploadResponse) error {
	return nil
}

func (s ScriptJsonRpcRouter) ShowLog(notify bool, request *model.ScriptUploadRequest, response *model.ScriptUploadResponse) error {
	return nil
}

func (s ScriptJsonRpcRouter) Exec(notify bool, request *model.ScriptExecRequest, response *model.ScriptExecResponse) error {
	return nil
}

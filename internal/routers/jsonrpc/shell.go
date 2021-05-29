package jsonrpc

import (
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/shell"
)

type ShellJsonRpcRouter struct {
}

func (s ShellJsonRpcRouter) Exec(notify bool, request *model.ShellExecRequest, response *model.ShellExecResponse) error {
	if notify {
		go func() {
			if err := s.Exec(false, request, response); err != nil {
				global.Logger.Warnf("执行Shell失败，失败原因:%+v", err)
			}
		}()
		return nil
	}

	if err := validate.Struct(request); err != nil {
		response.Error = *errcode.InvalidParams
		return nil
	}

	if request.Async {
		err := shell.AsyncExecuteShell(request.Command, global.AgentConfig.ExecLogPath, request.Logfile)
		if err != nil {
			response.Error = *errcode.NewBizErrorWithErr(err)
			return nil
		} else {
			response.Error = *errcode.Success
			return nil
		}
	} else {
		log, err := shell.ExecuteShell(request.Command)
		if err != nil {
			response.Error = *errcode.NewBizErrorWithErr(err)
			return nil
		} else {
			response.Log = log
			response.Error = *errcode.Success
			return nil
		}
	}
}

func (s ShellJsonRpcRouter) ShowLog(notify bool, request *model.ShellLogRequest, response *model.ShellLogResponse) error {
	return nil
}

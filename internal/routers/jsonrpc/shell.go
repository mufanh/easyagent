package jsonrpc

import (
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/convert"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/shell"
	"github.com/pkg/errors"
)

type CommandJsonRpcRouter struct {
}

func (s CommandJsonRpcRouter) Exec(notify bool, request *model.CommandExecRequest, response *model.CommandExecResponse) error {
	if notify {
		go func() {
			if err := s.Exec(false, request, response); err != nil {
				global.Logger.Warnf("执行Shell失败，失败原因:%+v", err)
			}
		}()
		return nil
	}

	if request.Async {
		if err := shell.AsyncExecuteShell(request.Command, global.AgentConfig.ExecLogPath, request.Logfile); err != nil {
			response.SetBizErr(errors.Wrap(err, "异步执行Shell失败"))
			return nil
		} else {
			response.SetErr(errcode.Success)
		}
	} else {
		if log, err := shell.ExecuteShell(request.Command, global.AgentConfig.ExecTimeout); err != nil {
			response.SetBizErr(errors.Wrap(err, "同步执行Shell失败"))
		} else {
			response.Log = convert.MustToCharsetStr(log, global.AgentConfig.Charset, "UTF-8")
			response.SetErr(errcode.Success)
		}
	}

	return nil
}

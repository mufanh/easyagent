package jsonrpc

import (
	"github.com/issue9/jsonrpc"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/shell"
	"github.com/mufanh/easyagent/pkg/util/fileutil"
	"github.com/pkg/errors"
	"path/filepath"
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

	if request.Async {
		if err := shell.AsyncExecuteShell(request.Command, global.AgentConfig.ExecLogPath, request.Logfile); err != nil {
			response.SetBizErr(errors.Wrap(err, "异步执行Shell失败"))
			return nil
		}
	} else {
		if log, err := shell.ExecuteShell(request.Command, global.AgentConfig.ExecTimeout); err != nil {
			response.SetBizErr(errors.Wrap(err, "同步执行Shell失败"))
			return nil
		} else {
			response.Log = log
		}
	}

	response.SetErr(errcode.Success)
	return nil
}

func (s ShellJsonRpcRouter) ShowLog(notify bool, request *model.ShellLogRequest, response *model.ShellLogResponse) error {
	if notify {
		return jsonrpc.NewError(jsonrpc.CodeInvalidRequest, "查看日志不能是通知型调用")
	}

	if bytes, err := fileutil.Read(filepath.Join(global.AgentConfig.ExecLogPath, request.Logfile)); err != nil {
		response.SetBizErr(errors.Wrap(err, "读取日志文件失败"))
		return nil
	} else {
		response.SetErr(errcode.Success)
		response.Log = string(bytes)
	}
	return nil
}

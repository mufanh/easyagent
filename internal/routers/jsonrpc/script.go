package jsonrpc

import (
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/shell"
	"github.com/mufanh/easyagent/pkg/util/fileutil"
	"github.com/pkg/errors"
	"path/filepath"
)

type ScriptJsonRpcRouter struct {
}

func (s ScriptJsonRpcRouter) Upload(notify bool, request *model.ScriptUploadRequest, response *model.ScriptUploadResponse) error {
	if notify {
		go func() {
			if err := s.Upload(false, request, response); err != nil {
				global.Logger.Warnf("脚本上送失败，错误原因:%+v", err)
			}
		}()
		return nil
	}

	filename := filepath.Join(global.AgentConfig.ScriptPath, request.GroupDir, request.Name)
	if exist, _ := fileutil.Exists(filename); exist {
		response.SetErr(errcode.NewBizErrorWithMsg("脚本已经存在"))
		return nil
	}

	if err := fileutil.MkdirAll(filepath.Join(global.AgentConfig.ScriptPath, request.GroupDir), 0700); err != nil {
		response.SetBizErr(errors.Wrap(err, "脚本上送目录创建失败"))
		return nil
	}

	if err := fileutil.Write(filename, []byte(request.Content)); err != nil {
		response.SetBizErr(errors.Wrap(err, "脚本写入失败"))
		return nil
	}

	response.SetErr(errcode.Success)
	return nil
}

func (s ScriptJsonRpcRouter) ShowLog(notify bool, request *model.ScriptUploadRequest, response *model.ScriptUploadResponse) error {
	return nil
}

func (s ScriptJsonRpcRouter) Exec(notify bool, request *model.ScriptExecRequest, response *model.ScriptExecResponse) error {
	if notify {
		go func() {
			if err := s.Exec(false, request, response); err != nil {
				global.Logger.Warnf("执行脚本失败，错误原因:%+v", err)
			}
		}()
		return nil
	}

	filename := filepath.Join(global.AgentConfig.ScriptPath, request.GroupDir, request.Name)
	if exist, _ := fileutil.Exists(filename); !exist {
		response.SetErr(errcode.NewBizErrorWithMsg("脚本不存在"))
		return nil
	}

	if request.Async {
		if err := shell.AsyncExecuteScript(filename, global.AgentConfig.ExecLogPath, request.Logfile); err != nil {
			response.SetBizErr(errors.Wrap(err, "异步执行Shell脚本失败"))
			return nil
		}
	} else {
		if log, err := shell.ExecuteScript(filename, global.AgentConfig.ExecTimeout); err != nil {
			response.SetBizErr(errors.Wrap(err, "同步执行Shell脚本失败"))
			return nil
		} else {
			response.Log = log
		}
	}

	response.SetErr(errcode.Success)
	return nil
}

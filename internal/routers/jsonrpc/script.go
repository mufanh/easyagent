package jsonrpc

import (
	"github.com/issue9/jsonrpc"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/convert"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/shell"
	"github.com/mufanh/easyagent/pkg/util/fileutil"
	"github.com/pkg/errors"
	"io/ioutil"
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

func (s ScriptJsonRpcRouter) Show(notify bool, request *model.ShowScriptRequest, response *model.ShowScriptResponse) error {
	if notify {
		return jsonrpc.NewError(jsonrpc.CodeInvalidRequest, "查看脚本不能是通知型服务")
	}

	filename := filepath.Join(global.AgentConfig.ScriptPath, request.GroupDir, request.Name)
	if exist, _ := fileutil.Exists(filename); !exist {
		response.SetErr(errcode.NewBizErrorWithMsg("脚本不存在"))
		return nil
	}

	if content, err := fileutil.Read(filename); err != nil {
		response.SetErr(errcode.NewBizErrorWithMsg("脚本读取失败"))
	} else {
		response.Content = string(content)
		response.SetErr(errcode.Success)
	}

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
		} else {
			response.SetErr(errcode.Success)
		}
	} else {
		if log, err := shell.ExecuteScript(filename, global.AgentConfig.ExecTimeout); err != nil {
			response.SetBizErr(errors.Wrap(err, "同步执行Shell脚本失败"))
		} else {
			response.Log = convert.MustToCharsetStr(log, global.AgentConfig.Charset, "UTF-8")
			response.SetErr(errcode.Success)
		}
	}

	return nil
}

func (s ScriptJsonRpcRouter) ShowGroupDirs(notify bool, request *model.ScriptShowGroupDirsRequest, response *model.ScriptShowGroupDirsResponse) error {
	if notify {
		return jsonrpc.NewError(jsonrpc.CodeInvalidRequest, "查看脚本分组目录不能是通知型服务")
	}

	if fileInfos, err := ioutil.ReadDir(global.AgentConfig.ScriptPath); err != nil {
		response.SetBizErr(errors.Wrap(err, "获取脚本分组目录列表失败"))
	} else {
		groupDirs := make([]string, 0)
		for i := 0; i < len(fileInfos); i++ {
			if fileInfos[i].IsDir() {
				groupDirs = append(groupDirs, fileInfos[i].Name())
			}
		}
		response.GroupDirs = groupDirs
		response.SetErr(errcode.Success)
	}

	return nil
}

func (s ScriptJsonRpcRouter) ShowScriptFiles(notify bool, request *model.ScriptShowFilesRequest, response *model.ScriptShowFilesResponse) error {
	if notify {
		return jsonrpc.NewError(jsonrpc.CodeInvalidRequest, "查看固定分组下脚本服务不能是通知型服务")
	}

	groupDir := filepath.Join(global.AgentConfig.ScriptPath, request.GroupDir)
	if fileInfos, err := ioutil.ReadDir(groupDir); err != nil {
		response.SetBizErr(errors.Wrap(err, "获取脚本分组下脚本列表失败"))
	} else {
		files := make([]string, 0)
		for i := 0; i < len(fileInfos); i++ {
			if !fileInfos[i].IsDir() {
				files = append(files, fileInfos[i].Name())
			}
		}
		response.ScriptFiles = files
		response.SetErr(errcode.Success)
	}

	return nil
}

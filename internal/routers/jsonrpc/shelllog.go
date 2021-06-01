package jsonrpc

import (
	"github.com/issue9/jsonrpc"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/convert"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/util/fileutil"
	"github.com/pkg/errors"
	"path/filepath"
)

type ShellLogJsonRpcRouter struct {
}

func (s ShellLogJsonRpcRouter) Show(notify bool, request *model.ShowShellLogRequest, response *model.ShowShellLogResponse) error {
	if notify {
		return jsonrpc.NewError(jsonrpc.CodeInvalidRequest, "查看日志不能是通知型调用")
	}

	if bytes, err := fileutil.Read(filepath.Join(global.AgentConfig.ExecLogPath, request.Logfile)); err != nil {
		response.SetBizErr(errors.Wrap(err, "读取日志文件失败"))
	} else {
		response.SetErr(errcode.Success)
		response.Log = convert.MustToCharsetStr(string(bytes), global.AgentConfig.Charset, "UTF-8")
	}
	return nil
}

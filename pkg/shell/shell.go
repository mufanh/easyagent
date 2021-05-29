package shell

import (
	"github.com/mufanh/easyagent/pkg/util/fileutil"
	"github.com/pkg/errors"
)

func ExecuteScript(filename string, timeout int) (string, error) {
	if content, err := fileutil.Read(filename); err != nil {
		return "", errors.Wrap(err, "执行脚本文件失败")
	} else {
		return ExecuteShell(string(content), timeout)
	}
}

func AsyncExecuteScript(filename string, logDir string, logFile string) error {
	if content, err := fileutil.Read(filename); err != nil {
		return errors.Wrap(err, "执行脚本文件失败")
	} else {
		return AsyncExecuteShell(string(content), logDir, logFile)
	}
}

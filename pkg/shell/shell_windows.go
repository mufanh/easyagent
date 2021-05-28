// +build windows

package shell

import (
	"bufio"
	"bytes"
	"github.com/mufanh/easyagent/global"
	"github.com/pkg/errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func ExecuteShell(command string) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("cmd.exe", "/C", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func AsyncExecuteShell(command string, tmpdir string, logfile string) error {
	err := os.MkdirAll(tmpdir, 0700)
	if err != nil {
		return errors.Wrap(err, "执行命令日志文件目录不存在，且自动创建失败")
	}

	file, err := os.Create(filepath.Join(tmpdir, "", logfile))
	if err != nil {
		return errors.Wrap(err, "执行命令日志文件创建失败")
	}
	defer func() {
		if err = file.Close(); err != nil {
			global.Logger.Warnf("关闭文件失败", err)
		}
	}()

	cmd := exec.Command("cmd.exe", "/C", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	stdoutPipe, err := cmd.StdoutPipe()
	writer := bufio.NewWriter(file)
	defer func() {
		if err = writer.Flush(); err != nil {
			global.Logger.Warnf("刷新执行命名日志文件失败", err)
		}
	}()

	err = cmd.Start()
	go func() {
		if _, err = io.Copy(writer, stdoutPipe); err != nil {
			global.Logger.Warnf("写入执行命名日志文件失败", err)
		}
	}()

	return nil
}

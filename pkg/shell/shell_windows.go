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
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func AsyncExecuteShell(command string, logDir string, logFile string) error {
	if err := os.MkdirAll(logDir, 0700); err != nil {
		return errors.Wrap(err, "执行命令日志文件目录不存在，且自动创建失败")
	}

	file, err := os.Create(filepath.Join(logDir, "", logFile))
	if err != nil {
		return errors.Wrap(err, "执行命令日志文件创建失败")
	}

	cmd := exec.Command("cmd.exe", "/C", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		if err = file.Close(); err != nil {
			global.Logger.Warnf("关闭文件失败", err)
		}
		return errors.Wrap(err, "执行命令日志输出管道创建失败")
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		if err = file.Close(); err != nil {
			global.Logger.Warnf("关闭文件失败", err)
		}
		return errors.Wrap(err, "执行命令日志输出管道创建失败")
	}

	go func() {
		defer func() {
			if err = file.Close(); err != nil {
				global.Logger.Warnf("关闭文件失败", err)
			}
		}()

		writer := bufio.NewWriter(file)
		if err = cmd.Start(); err != nil {
			global.Logger.Warnf("异步执行命令失败", err)
		}

		if _, err = io.Copy(writer, stdoutPipe); err != nil {
			global.Logger.Warnf("写入执行命名日志文件失败", err)
		}
		if _, err = io.Copy(writer, stderrPipe); err != nil {
			global.Logger.Warnf("写入执行命名日志文件失败", err)
		}
	}()

	return nil
}

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
	"time"
)

func asyncExecuteCommand(command string, params []string, logDir string, logFile string) error {
	if err := os.MkdirAll(logDir, 0700); err != nil {
		return errors.Wrap(err, "执行命令日志文件目录不存在，且自动创建失败")
	}

	filename := filepath.Join(logDir, "", logFile)
	file, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err, "执行命令日志文件创建失败")
	}

	cmd := exec.Command(command, params...)

	stdout, err1 := cmd.StdoutPipe()
	if err1 != nil {
		_ = file.Close()
		return errors.Wrap(err1, "执行命令输出管道创建失败")
	}
	stderr, err2 := cmd.StderrPipe()
	if err2 != nil {
		_ = file.Close()
		return errors.Wrap(err2, "执行命令错误输出管道创建失败")
	}

	go func() {
		defer func() {
			_ = file.Close()
		}()

		done := make(chan error)

		writer := bufio.NewWriter(file)
		go func() {
			if _, err := io.Copy(writer, stderr); err != nil {
				global.Logger.Warnf("脚本执行错误日志写入失败，脚本%s，错误信息:%+v", filename, err)
				done <- err
			}
		}()

		if err := cmd.Start(); err != nil {
			global.Logger.Warnf("脚本执行异常，脚本%s，错误信息:%+v", filename, err)
			return
		}

		go func() {
			done <- cmd.Wait()
		}()

		go func() {
			if _, err := io.Copy(writer, stdout); err != nil {
				global.Logger.Warnf("脚本执行日志写入失败，脚本%s，错误信息:%+v", filename, err)
				done <- err
			}
		}()

		select {
		case <-done:
			return
		}
	}()

	return nil
}

func executeCommand(command string, params []string, timeout int) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command(command, params...)
	cmd.Stderr = &out
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		return "", err
	}

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	after := time.After(time.Duration(timeout) * time.Millisecond)
	select {
	case <-after:
		_ = cmd.Process.Signal(syscall.SIGINT)
		time.Sleep(10 * time.Millisecond)
		_ = cmd.Process.Kill()
		return "", errors.New("命令执行超时")
	case err := <-done:
		return out.String(), err
	}
}

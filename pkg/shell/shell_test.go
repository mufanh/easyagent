// +build windows

package shell

import (
	"github.com/mufanh/easyagent/pkg/convert"
	"testing"
	"time"
)

func TestExecuteShell(t *testing.T) {
	if log, err := executeCommand("cmd.exe", []string{"/C", "ping www.baidu.com"}, 1000000); err != nil {
		t.Errorf("执行命令失败，错误信息:%+v", err)
	} else {
		t.Logf("执行名称成功，输出:\n%s", convert.MustToCharsetStr(log, "GB18030", "UTF-8"))
	}
}

func TestExecuteScript(t *testing.T) {
	if log, err := executeCommand("cmd.exe", []string{"/C", "D:\\Workspace\\projects4\\easyagent\\_temp\\scripts\\test\\ping.bat"}, 1000000); err != nil {
		t.Errorf("执行命令失败，错误信息:%+v", err)
	} else {
		t.Logf("执行名称成功，输出:\n%s", convert.MustToCharsetStr(log, "GB18030", "UTF-8"))
	}
}

func TestAsyncExecuteShell(t *testing.T) {
	if err := asyncExecuteCommand("cmd.exe", []string{"/C", "ping www.baidu.com"}, "D:\\Workspace\\projects4\\easyagent\\_temp\\runlogs", "test.log"); err != nil {
		t.Errorf("执行命令失败，错误信息:%+v", err)
	}
	after := time.After(time.Duration(1) * time.Minute)
	select {
	case <-after:
		return
	}
}

func TestAsyncExecuteScript(t *testing.T) {
	if err := asyncExecuteCommand("cmd.exe", []string{"/C", "D:\\Workspace\\projects4\\easyagent\\_temp\\scripts\\test\\ping.bat"}, "D:\\Workspace\\projects4\\easyagent\\_temp\\runlogs", "test.log"); err != nil {
		t.Errorf("执行命令失败，错误信息:%+v", err)
	}
	after := time.After(time.Duration(1) * time.Minute)
	select {
	case <-after:
		return
	}
}

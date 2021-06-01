package main

import (
	"context"
	"github.com/gookit/goutil/strutil"
	"github.com/gorilla/websocket"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/internal/routers"
	"github.com/mufanh/easyagent/pkg/util/netutil"
	"github.com/nightlyone/lockfile"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func init() {
	if err := global.SetupAgentSetting("configs/", "agent"); err != nil {
		panic(errors.Wrap(err, "初始化应用配置失败"))
	}
	if err := global.SetupLogger(
		global.AgentLogConfig.LogFilepath,
		global.AgentLogConfig.LogFilename,
		global.AgentLogConfig.LogMaxSize,
		global.AgentLogConfig.LogMaxAge); err != nil {
		panic(errors.Wrap(err, "初始化应用日志失败"))
	}
}

func main() {
	if lock, err := lockfile.New(filepath.Join(os.TempDir(), strutil.Md5(global.AgentConfig.WsAddr))); err != nil {
		global.Logger.Fatalf("启动easyagent-agent失败，创建lock文件失败，详细错误原因:%+v", err)
		return
	} else {
		if err := lock.TryLock(); err != nil {
			global.Logger.Fatalf("启动easyagent-agent失败，获取lock文件失败，详细错误原因:%+v", err)
			return
		}
		defer func() {
			if err := lock.Unlock(); err != nil {
				global.Logger.Warnf("lock文件解锁失败，详细错误原因:%+v", err)
			}
		}()
	}

	requestHeader, err := prepareRequestHeader()
	if err != nil {
		global.Logger.Fatalf("初始化Websocket Header失败，启动应用失败，详细错误原因:%+v", err)
		return
	}

	conn, _, err := websocket.DefaultDialer.Dial(global.AgentConfig.WsAddr, *requestHeader)
	if err != nil {
		global.Logger.Fatalf("连接服务器地址%s失败，详细错误原因:%+v", global.AgentConfig.WsAddr, err)
		return
	}
	global.AgentRepo.SetConn(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := routers.NewAgentJsonRpcRouter()
	client := router.NewConn(global.AgentRepo.Transport(), nil)
	if err = client.Serve(ctx); err != nil {
		global.Logger.Fatalf("连接Websocket服务失败，应用启动失败，详细错误原因:%+v", err)
		return
	}
}

func prepareRequestHeader() (*http.Header, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "获取当前用户信息失败")
	}

	ips, err := netutil.GetLocalIPsStr()
	if err != nil {
		return nil, errors.Wrap(err, "获取本地IP列表失败")
	}

	agentInfo := model.AgentInfo{
		Token:       global.AgentConfig.Token,
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Gid:         currentUser.Gid,
		Uid:         currentUser.Uid,
		Username:    currentUser.Username,
		Name:        currentUser.Name,
		HomeDir:     currentUser.HomeDir,
		ScriptPath:  global.AgentConfig.ScriptPath,
		ExecLogPath: global.AgentConfig.ExecLogPath,
		Desc:        global.AgentConfig.Desc,
		IPList:      ips,
	}
	requestHeader := http.Header{}
	for k, v := range *model.ConvertAgentInfo2Map(&agentInfo) {
		requestHeader.Add(k, v)
	}

	return &requestHeader, nil
}

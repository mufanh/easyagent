package main

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/issue9/jsonrpc"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	jsonrpc2 "github.com/mufanh/easyagent/internal/routers/jsonrpc"
	"github.com/pkg/errors"
	"net/http"
	"os/user"
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
	router := prepareJsonRpcRouter()

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
	global.SetConn(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := router.NewConn(global.GetTransport(), nil)
	if err = client.Serve(ctx); err != nil {
		global.Logger.Fatalf("连接Websocket服务失败，应用启动失败，详细错误原因:%+v", err)
		return
	}
}

func prepareJsonRpcRouter() *jsonrpc.Server {
	server := new(jsonrpc.Server)
	server.Register("session.close", jsonrpc2.CloseSession)
	return server
}

func prepareRequestHeader() (*http.Header, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "获取当前用户信息失败")
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
	}
	requestHeader := http.Header{}
	for k, v := range *model.ConvertAgentInfo2Map(&agentInfo) {
		requestHeader.Add(k, v)
	}

	return &requestHeader, nil
}

package model

import "github.com/mufanh/easyagent/pkg/errcode"

type ScriptUploadRequest struct {
	Token   string `json:"token"`
	Content string `json:"content"`
	// 脚本分组名（脚本上一级目录名）
	GroupDir string `json:"groupDir"`
	// 脚本名
	Name string `json:"name"`
	// 脚本对应的操作系统
	OS string `json:"os"`
}

type ScriptUploadResponse struct {
	errcode.Error
}

type ScriptLogRequest struct {
	Token string `json:"token"`
	// 脚本分组名（脚本上一级目录名）
	GroupDir string `json:"groupDir"`
	// 脚本名
	Name string `json:"name"`
}

type ScriptLogResponse struct {
	errcode.Error
	// 日志
	Log string `json:"log"`
}

type ScriptExecRequest struct {
	Token string `json:"token"`
	// 脚本分组名（脚本上一级目录名）
	GroupDir string `json:"groupDir"`
	// 脚本名
	Name string `json:"name"`

	Async   bool   `json:"async"`
	Logfile string `json:"logfile"`
}

type ScriptExecResponse struct {
	errcode.Error
	// 若async=false，那么日志会直接记录到该字段返回
	Log string `json:"log"`
}
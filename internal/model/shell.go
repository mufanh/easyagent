package model

import "github.com/mufanh/easyagent/pkg/errcode"

type ShellExecRequest struct {
	Token   string `json:"token"`
	Command string `json:"command"`
	Async   bool   `json:"async"`
	Logfile string `json:"logfile"`
}

type ShellExecResponse struct {
	ShellExecRequest
	errcode.Error
	// 若async=false，那么日志会直接记录到该字段返回
	Log string `json:"log"`
}

func NewErrorShellExecResponse(request ShellExecRequest, err errcode.Error) *ShellExecResponse {
	response := ShellExecResponse{request, err, ""}
	return &response
}

func NewSuccessShellExecResponse(request ShellExecRequest, log string) *ShellExecResponse {
	response := ShellExecResponse{request, *errcode.Success, log}
	return &response

}

package model

type ShellExecRequest struct {
	Token   string `json:"token"`
	Command string `json:"command"`
	Async   bool   `json:"async"`
	Logfile string `json:"logfile"`
}

type ShellExecResponse struct {
	BaseResponse
	// 若async=false，那么日志会直接记录到该字段返回
	Log string `json:"log"`
}

type ShellLogRequest struct {
	Token   string `json:"token"`
	Logfile string `json:"logfile"`
}

type ShellLogResponse struct {
	BaseResponse
	// 日志
	Log string `json:"log"`
}

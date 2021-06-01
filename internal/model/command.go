package model

type CommandExecRequest struct {
	Token   string `json:"token"`
	Command string `json:"command"`
	Async   bool   `json:"async"`
	Logfile string `json:"logfile"`
}

type CommandExecResponse struct {
	BaseResponse
	// 若async=false，那么日志会直接记录到该字段返回
	Log string `json:"log,omitempty"`
}

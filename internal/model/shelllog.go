package model

type ShowShellLogRequest struct {
	Token   string `json:"token"`
	Logfile string `json:"logfile"`
}

type ShowShellLogResponse struct {
	BaseResponse
	// 日志
	Log string `json:"log"`
}

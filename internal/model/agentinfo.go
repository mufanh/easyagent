package model

import "github.com/pkg/errors"

const (
	tokenHttpHeaderKey       = "_AGENT_INFO_TOKEN"
	ConnectTimeHttpHeaderKey = "_AGENT_INFO_CONNECT_TIME"
	osHttpHeaderKey          = "_AGENT_INFO_OS"
	archHttpHeaderKey        = "_AGENT_INFO_ARCH"
	gidHttpHeaderKey         = "_AGENT_INFO_GID"
	uidHttpHeaderKey         = "_AGENT_INFO_UID"
	usernameHttpHeaderKey    = "_AGENT_INFO_USERNAME"
	nameHttpHeaderKey        = "_AGENT_INFO_NAME"
	homeDirHttpHeaderKey     = "_AGENT_INFO_HOME_DIR"
	scriptPathHttpHeaderKey  = "_AGENT_INFO_SCRIPT_PATH"
	execLogPathHttpHeaderKey = "_AGENT_INFO_EXEC_LOG_PATH"
	descHttpHeaderKey        = "_AGENT_INFO_DESC"
	charsetHttpHeaderKey     = "_AGENT_INFO_CHARSET"
	ipsHttpHeaderKey         = "_AGENT_INFO_IPS"
)

type AgentInfo struct {
	Token       string `json:"token"`
	ConnectTime string `json:"connect_time"`
	OS          string `json:"os"`
	Arch        string `json:"arch"`
	Gid         string `json:"gid"`
	Uid         string `json:"uid"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	HomeDir     string `json:"home_dir"`
	ScriptPath  string `json:"script_path"`
	ExecLogPath string `json:"exec_log_path"`
	Desc        string `json:"desc"`
	Charset     string `json:"charset"`
	IPList      string `json:"ips"`
}

func ConvertMap2AgentInfo(data map[string]string) *AgentInfo {
	return &AgentInfo{
		Token:       data[tokenHttpHeaderKey],
		ConnectTime: data[ConnectTimeHttpHeaderKey],
		OS:          data[osHttpHeaderKey],
		Arch:        data[archHttpHeaderKey],
		Gid:         data[gidHttpHeaderKey],
		Uid:         data[uidHttpHeaderKey],
		Username:    data[usernameHttpHeaderKey],
		Name:        data[nameHttpHeaderKey],
		HomeDir:     data[homeDirHttpHeaderKey],
		ScriptPath:  data[scriptPathHttpHeaderKey],
		ExecLogPath: data[execLogPathHttpHeaderKey],
		Desc:        data[descHttpHeaderKey],
		Charset:     data[charsetHttpHeaderKey],
		IPList:      data[ipsHttpHeaderKey],
	}
}

func ConvertAgentInfo2Map(agentInfo *AgentInfo) *map[string]string {
	r := make(map[string]string)
	r[tokenHttpHeaderKey] = agentInfo.Token
	r[ConnectTimeHttpHeaderKey] = agentInfo.ConnectTime
	r[osHttpHeaderKey] = agentInfo.OS
	r[archHttpHeaderKey] = agentInfo.Arch
	r[gidHttpHeaderKey] = agentInfo.Gid
	r[uidHttpHeaderKey] = agentInfo.Uid
	r[usernameHttpHeaderKey] = agentInfo.Username
	r[nameHttpHeaderKey] = agentInfo.Name
	r[homeDirHttpHeaderKey] = agentInfo.HomeDir
	r[scriptPathHttpHeaderKey] = agentInfo.ScriptPath
	r[execLogPathHttpHeaderKey] = agentInfo.ExecLogPath
	r[descHttpHeaderKey] = agentInfo.Desc
	r[charsetHttpHeaderKey] = agentInfo.Charset
	r[ipsHttpHeaderKey] = agentInfo.IPList
	return &r
}

func (s *AgentInfo) Check() error {
	if s.Token == "" {
		return errors.New("Token不能为空")
	}
	return nil
}

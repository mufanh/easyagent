package model

import "github.com/pkg/errors"

var (
	TokenHttpHeaderKey       = "_AGENT_INFO_TOKEN"
	OSHttpHeaderKey          = "_AGENT_INFO_OS"
	ArchHttpHeaderKey        = "_AGENT_INFO_ARCH"
	GidHttpHeaderKey         = "_AGENT_INFO_GID"
	UidHttpHeaderKey         = "_AGENT_INFO_UID"
	UsernameHttpHeaderKey    = "_AGENT_INFO_USERNAME"
	NameHttpHeaderKey        = "_AGENT_INFO_NAME"
	HomeDirHttpHeaderKey     = "_AGENT_INFO_HOME_DIR"
	ScriptPathHttpHeaderKey  = "_AGENT_INFO_SCRIPT_PATH"
	ExecLogPathHttpHeaderKey = "_AGENT_INFO_EXEC_LOG_PATH"
	DescHttpHeaderKey        = "_AGENT_INFO_DESC"
)

type AgentInfo struct {
	Token       string `json:"token"`
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
}

func ConvertMap2AgentInfo(data map[string]string) *AgentInfo {
	return &AgentInfo{
		Token:       data[TokenHttpHeaderKey],
		OS:          data[OSHttpHeaderKey],
		Arch:        data[ArchHttpHeaderKey],
		Gid:         data[GidHttpHeaderKey],
		Uid:         data[UidHttpHeaderKey],
		Username:    data[UsernameHttpHeaderKey],
		Name:        data[NameHttpHeaderKey],
		HomeDir:     data[HomeDirHttpHeaderKey],
		ScriptPath:  data[ScriptPathHttpHeaderKey],
		ExecLogPath: data[ExecLogPathHttpHeaderKey],
		Desc:        data[DescHttpHeaderKey],
	}
}

func ConvertAgentInfo2Map(agentInfo *AgentInfo) *map[string]string {
	r := make(map[string]string)
	r[TokenHttpHeaderKey] = agentInfo.Token
	r[OSHttpHeaderKey] = agentInfo.OS
	r[ArchHttpHeaderKey] = agentInfo.Arch
	r[GidHttpHeaderKey] = agentInfo.Gid
	r[UidHttpHeaderKey] = agentInfo.Uid
	r[UsernameHttpHeaderKey] = agentInfo.Username
	r[NameHttpHeaderKey] = agentInfo.Name
	r[HomeDirHttpHeaderKey] = agentInfo.HomeDir
	r[ScriptPathHttpHeaderKey] = agentInfo.ScriptPath
	r[ExecLogPathHttpHeaderKey] = agentInfo.ExecLogPath
	r[DescHttpHeaderKey] = agentInfo.Desc
	return &r
}

func (s *AgentInfo) Check() error {
	if s.Token == "" {
		return errors.New("Token不能为空")
	}
	return nil
}

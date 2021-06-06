package global

import (
	"flag"
	"github.com/mufanh/easyagent/pkg/setting"
)

var (
	AgentConfig    *AgentSettings
	AgentLogConfig *LogSettings

	// 允许命令行修改的地址
	wsAddr  string
	charset string
)

func init() {
	// 设置允许从命令行参数获取的参数值
	flag.StringVar(&wsAddr, "wsAddr", "", "Server的websocket服务")
	// 设置字符集
	flag.StringVar(&charset, "charset", "", "字符集")
	flag.Parse()
}

func SetupAgentSetting(configPath string, configName string) error {
	settings, err := setting.NewSetting(configPath, configName)
	if err != nil {
		return err
	}

	if err = settings.ReadSection("Agent", &AgentConfig); err != nil {
		return err
	}

	// 命令行优先级更高，目前只允许命令行修改websocket地址和字符集
	if wsAddr != "" {
		AgentConfig.WsAddr = wsAddr
	}
	if charset != "" {
		AgentConfig.Charset = charset
	}

	if err = AgentConfig.check(); err != nil {
		return err
	}

	if err = settings.ReadSection("Log", &AgentLogConfig); err != nil {
		return err
	}
	if err = AgentLogConfig.check(); err != nil {
		return err
	}
	return nil
}

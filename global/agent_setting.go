package global

import (
	"github.com/mufanh/easyagent/pkg/setting"
)

var (
	AgentConfig    *AgentSettings
	AgentLogConfig *LogSettings
)

func SetupAgentSetting(configPath string, configName string) error {
	settings, err := setting.NewSetting(configPath, configName)
	if err != nil {
		return err
	}

	if err = settings.ReadSection("Agent", &AgentConfig); err != nil {
		return err
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

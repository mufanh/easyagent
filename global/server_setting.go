package global

import "github.com/mufanh/easyagent/pkg/setting"

var (
	ServerConfig    *ServerSettings
	ServerLogConfig *LogSettings
)

func SetupServerSetting(configPath string, configName string) error {
	settings, err := setting.NewSetting(configPath, configName)
	if err != nil {
		return err
	}

	if err = settings.ReadSection("Server", &ServerConfig); err != nil {
		return err
	}
	if err = ServerConfig.check(); err != nil {
		return err
	}

	if err = settings.ReadSection("Log", &ServerLogConfig); err != nil {
		return err
	}
	if err = ServerLogConfig.check(); err != nil {
		return err
	}
	return nil
}

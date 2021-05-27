package setting

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configPath string, configName string) (*Setting, error) {
	if configName == "" {
		return nil, errors.New("配置文件名不能为空")
	}
	if configPath == "" {
		configPath = "configs/"
	}
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AddConfigPath(configPath)
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "读取配置文件 "+configName+" 失败")
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return errors.Wrap(err, "解析 "+k+" 失败")
	}
	return nil
}

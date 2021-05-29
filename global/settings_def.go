package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"time"
)

var validate = validator.New()

type ServerSettings struct {
	// debug/release/test
	RunMode string

	HttpPort     uint16 `validate:"gte=5001,lte=65535"`
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	DefaultPageSize int
	MaxPageSize     int
}

type AgentSettings struct {
	WsAddr      string
	Token       string
	TmpPath     string
	ScriptPath  string
	ExecLogPath string
	Desc        string
	ExecTimeout int
}

type LogSettings struct {
	LogFilepath string
	LogFilename string
	LogMaxSize  int
	LogMaxAge   int
}

func (s *ServerSettings) check() error {
	if err := validate.Struct(s); err != nil {
		return errors.Wrap(err, "校验ServerSettings配置不合法")
	}
	return nil
}

func (s *AgentSettings) check() error {
	if err := validate.Struct(s); err != nil {
		return errors.Wrap(err, "校验AgentSettings配置不合法")
	}
	return nil
}

func (s *LogSettings) check() error {
	if err := validate.Struct(s); err != nil {
		return errors.Wrap(err, "校验LogSettings配置不合法")
	}
	return nil
}

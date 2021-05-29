package jsonrpc

import (
	"github.com/go-playground/validator/v10"
	"github.com/issue9/jsonrpc"
	"github.com/mufanh/easyagent/global"
	"github.com/pkg/errors"
)

var validate = validator.New()

type SessionJsonRpcRouter struct {
}

func (s SessionJsonRpcRouter) Close(notify bool, request *interface{}, response *interface{}) error {
	if notify {
		go func() {
			if err := s.Close(false, request, response); err != nil {
				global.Logger.Warnf("关闭连接失败，错误原因:%+v", err)
			}
		}()
		return nil
	}

	if err := validate.Struct(request); err != nil {
		return jsonrpc.NewErrorWithError(jsonrpc.CodeInvalidRequest, err)
	}

	if conn := global.AgentRepo.Conn(); conn != nil {
		if err := conn.Close(); err != nil {
			return jsonrpc.NewErrorWithError(jsonrpc.CodeInternalError, errors.Wrap(err, "连接关闭失败"))
		}
	} else {
		return jsonrpc.NewErrorWithError(jsonrpc.CodeInternalError, errors.New("连接不存在"))
	}

	return nil
}

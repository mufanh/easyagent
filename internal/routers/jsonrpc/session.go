package jsonrpc

import (
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/pkg/errors"
)

type SessionJsonRpcRouter struct {
}

func (s SessionJsonRpcRouter) Close(notify bool, request *interface{}, response *model.SessionCloseResponse) error {
	if notify {
		go func() {
			if err := s.Close(false, request, response); err != nil {
				global.Logger.Warnf("关闭连接失败，错误原因:%+v", err)
			}
		}()
		return nil
	}

	if conn := global.AgentRepo.Conn(); conn != nil {
		if err := conn.Close(); err != nil {
			response.SetBizErr(errors.Wrap(err, "连接关闭失败"))
			return nil
		}
		return nil
	} else {
		response.SetBizErr(errcode.NewBizErrorWithMsg("连接不存在"))
		return nil
	}
}

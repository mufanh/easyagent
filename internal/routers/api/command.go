package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/result"
	"io/ioutil"
)

type CommandApiRouter struct {
}

func (s CommandApiRouter) Exec(c *gin.Context) {
	responseWriter := result.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.CommandExecRequest
	if err = json.Unmarshal(body, &request); err != nil {
		responseWriter.ToErrorResponse(errcode.InvalidParams)
		return
	}

	conn := global.ServerRepo.SessionJConn(request.Token)
	if conn == nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithMsg("Token不存在"))
		return
	}

	done := make(chan bool)
	if err := conn.Send("shell.exec", &request, func(response *model.CommandExecResponse) error {
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

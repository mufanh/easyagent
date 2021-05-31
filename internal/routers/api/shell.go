package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/mufanh/easyagent/pkg/app"
	"github.com/mufanh/easyagent/pkg/errcode"
	"io/ioutil"
)

type ShellApiRouter struct {
}

func (s ShellApiRouter) Exec(c *gin.Context) {
	responseWriter := app.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ShellExecRequest
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
	if err := conn.Send("shell.exec", &request, func(response *model.ShellExecResponse) error {
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ShellApiRouter) ShowLog(c *gin.Context) {
	responseWriter := app.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ShellLogRequest
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
	if err := conn.Send("shell.log", &request, func(response *model.ShellLogResponse) error {
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

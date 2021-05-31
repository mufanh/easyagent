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

type ScriptApiRouter struct {
}

func (s ScriptApiRouter) Upload(c *gin.Context) {
	responseWriter := app.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ScriptUploadRequest
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
	if err := conn.Send("script.upload", &request, func(response *model.ScriptUploadResponse) error {
		responseWriter.ToSuccessResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ScriptApiRouter) ShowLog(c *gin.Context) {
	responseWriter := app.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ScriptLogRequest
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
	if err := conn.Send("script.log", &request, func(response *model.ScriptLogResponse) error {
		responseWriter.ToSuccessResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ScriptApiRouter) Exec(c *gin.Context) {
	responseWriter := app.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ScriptExecRequest
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
	if err := conn.Send("script.exec", &request, func(response *model.ScriptExecResponse) error {
		if response.IsSuccess() {
			responseWriter.ToSuccessResponse(response)
		} else {
			responseWriter.ToErrorResponse(&response.Error)
		}
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

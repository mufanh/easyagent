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

type ScriptApiRouter struct {
}

func (s ScriptApiRouter) Upload(c *gin.Context) {
	responseWriter := result.NewResponse(c)

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
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ScriptApiRouter) ShowLog(c *gin.Context) {
	responseWriter := result.NewResponse(c)

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
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ScriptApiRouter) Exec(c *gin.Context) {
	responseWriter := result.NewResponse(c)

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
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ScriptApiRouter) ShowGroupDirs(c *gin.Context) {
	responseWriter := result.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ScriptShowGroupDirsRequest
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
	if err := conn.Send("script.showGroupDirs", &request, func(response *model.ScriptShowGroupDirsResponse) error {
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ScriptApiRouter) ShowScriptFiles(c *gin.Context) {
	responseWriter := result.NewResponse(c)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responseWriter.ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ScriptShowFilesRequest
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
	if err := conn.Send("script.showScriptFiles", &request, func(response *model.ScriptShowFilesResponse) error {
		responseWriter.ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

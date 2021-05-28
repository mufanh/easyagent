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

var (
	ErrorTokenNotFound   = errcode.NewError(30030001, "TOKEN不存在")
	ErrorShellExecFailed = errcode.NewError(30030002, "Shell执行失败")
)

type ShellApiRouter struct {
}

// @Summary
// @Produce json
// @Success 200 {object} model.AgentInfo "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/shell/exec [post]
func (s ShellApiRouter) Exec(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ShellExecRequest
	if err = json.Unmarshal(body, &request); err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
		return
	}

	conn := global.ServerRepo.GetSessionJConn(request.Token)
	if conn == nil {
		app.NewResponse(c).ToErrorResponse(ErrorTokenNotFound)
		return
	}

	var r *model.ShellExecResponse
	var e *errcode.Error
	done := make(chan bool)
	if err := conn.Send("shell.exec", &request, func(response *model.ShellExecResponse) error {
		r = response
		done <- true
		return nil
	}); err != nil {
		done <- true
		e = ErrorShellExecFailed
	}
	<-done
	if e != nil {
		app.NewResponse(c).ToSuccessResponse(e)
	} else {
		app.NewResponse(c).ToSuccessResponse(r)
	}
}

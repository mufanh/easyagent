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

// @Summary
// @Produce json
// @Success 200 {object}  "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/script/upload [post]
func (s ScriptApiRouter) Upload(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.ServerError)
		return
	}

	var request model.ScriptUploadRequest
	if err = json.Unmarshal(body, &request); err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
		return
	}

	conn := global.ServerRepo.SessionJConn(request.Token)
	if conn == nil {
		app.NewResponse(c).ToErrorResponse(errcode.NewBizErrorWithMsg("Token不存在"))
		return
	}

	done := make(chan bool)
	if err := conn.Send("script.upload", &request, func(response *model.ScriptUploadResponse) error {
		app.NewResponse(c).ToResponse(response)
		done <- true
		return nil
	}); err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.NewBizErrorWithErr(err))
		done <- false
	}
	<-done
}

func (s ScriptApiRouter) ShowLog(c *gin.Context) {

}

func (s ScriptApiRouter) Exec(c *gin.Context) {

}

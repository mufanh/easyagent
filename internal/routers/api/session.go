package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/pkg/app"
	"github.com/mufanh/easyagent/pkg/errcode"
)

var (
	ErrorCloseSession = errcode.NewError(20020001, "关闭连接失败")
)

type SessionApiRouter struct {
}

// @Summary 获取agent连接列表
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.AgentInfo "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/sessions [get]
func (s SessionApiRouter) List(c *gin.Context) {
	agentInfos := global.ServerRepo.ListAgentInfos()
	app.NewResponse(c).ToResponse(agentInfos)
}

// @Summary 关闭agent连接
// @Produce json
// @Param token param string "" "连接标识"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/sessions/{id} [delete]
func (s SessionApiRouter) Close(c *gin.Context) {
	token := c.Param("token")
	if err := global.ServerRepo.DeleteSession(token); err != nil {
		app.NewResponse(c).ToErrorResponse(ErrorCloseSession)
	} else {
		app.NewResponse(c).ToErrorResponse(errcode.Success)
	}

}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/pkg/app"
	"github.com/mufanh/easyagent/pkg/errcode"
)

type SessionApiRouter struct {
}

func (s SessionApiRouter) List(c *gin.Context) {
	agentInfos := global.ServerRepo.AgentInfos()
	app.NewResponse(c).ToSuccessResponse(agentInfos)
}

func (s SessionApiRouter) Close(c *gin.Context) {
	responseWriter := app.NewResponse(c)

	token := c.Param("token")
	if err := global.ServerRepo.DeleteSession(token); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
	} else {
		responseWriter.ToSuccessResponse(errcode.Success)
	}
}

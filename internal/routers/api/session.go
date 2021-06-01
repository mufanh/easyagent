package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/result"
)

type SessionApiRouter struct {
}

func (s SessionApiRouter) List(c *gin.Context) {
	agentInfos := global.ServerRepo.AgentInfos()
	result.NewResponse(c).ToResponse(agentInfos)
}

func (s SessionApiRouter) Close(c *gin.Context) {
	responseWriter := result.NewResponse(c)

	token := c.Param("token")
	if err := global.ServerRepo.DeleteSession(token); err != nil {
		responseWriter.ToErrorResponse(errcode.NewBizErrorWithErr(err))
	} else {
		responseWriter.ToResponse(errcode.Success)
	}
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/pkg/result"
)

type SessionApiRouter struct {
}

func (s SessionApiRouter) List(c *gin.Context) {
	agentInfos := global.ServerRepo.AgentInfos()
	result.NewResponse(c).ToResponse(agentInfos)
}

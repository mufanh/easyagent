package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/pkg/errcode"
	"github.com/mufanh/easyagent/pkg/result"
)

func Recovery() func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %+v", err)

				result.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
	}
}

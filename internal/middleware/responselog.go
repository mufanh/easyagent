package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/mufanh/easyagent/global"
	"github.com/mufanh/easyagent/pkg/logger"
	"time"
)

type ResponseLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func ResponseLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &ResponseLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{"response": bodyWriter.body.String()}
		global.Logger.WithFields(fields).Infof("访问日志: 方法:%s, 状态码:%d, 开始时间:%d, 结束时间:%d",
			c.Request.Method, bodyWriter.Status(), beginTime, endTime)
	}
}

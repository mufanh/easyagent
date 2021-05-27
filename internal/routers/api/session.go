package api

import "github.com/gin-gonic/gin"

type Session struct {
}

// @Summary 获取agent连接列表
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.SessionInfo "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/sessions [get]
func (s Session) List(c *gin.Context) {}

// @Summary 关闭agent连接
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/sessions/{id} [delete]
func (s Session) Close(c *gin.Context) {}

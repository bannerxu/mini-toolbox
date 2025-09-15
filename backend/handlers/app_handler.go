package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppHandler 应用程序处理器
type AppHandler struct{}

// NewAppHandler 创建新的应用程序处理器
func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

// HomePage 首页处理函数
func (h *AppHandler) HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "欢迎使用 Mini Toolbox Web 应用！",
		"version": "1.0.0",
		"status":  "运行中",
	})
}

// HealthCheck 健康检查端点
func (h *AppHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "健康",
		"timestamp": gin.H{
			"server": "运行正常",
		},
	})
}

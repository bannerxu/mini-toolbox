package main

import (
	"log"
	"os"

	"mini-toolbox/models"
	"mini-toolbox/routes"
)

func main() {
	// 创建必要的目录
	uploadDir := "uploads"
	compressedDir := "compressed"

	// 确保目录存在
	os.MkdirAll(uploadDir, 0755)
	os.MkdirAll(compressedDir, 0755)

	// 创建用户服务
	userService := models.NewInMemoryUserService()

	// 设置路由
	r := routes.SetupRoutes(userService)

	// 启动服务器在8080端口（与前端配置保持一致）
	port := ":8080"
	log.Printf("Mini Toolbox 后端服务启动在端口 %s", port)
	log.Printf("访问 http://localhost%s 查看应用", port)
	log.Printf("访问 http://localhost%s/api/v1/users 查看用户API", port)
	log.Printf("访问 http://localhost%s/api/v1/images 查看图片API", port)

	if err := r.Run(port); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}

package routes

import (
	"mini-toolbox/handlers"
	"mini-toolbox/models"
	"os"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置应用程序路由
func SetupRoutes(userService models.UserService) *gin.Engine {
	// 创建 Gin 路由器
	r := gin.Default()

	// 添加全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 创建处理器
	appHandler := handlers.NewAppHandler()
	userHandler := handlers.NewUserHandler(userService)

	// 创建图片服务和处理器
	uploadDir := "uploads"
	compressedDir := "compressed"

	// 确保目录存在
	os.MkdirAll(uploadDir, 0755)
	os.MkdirAll(compressedDir, 0755)

	imageService := models.NewDefaultImageService(uploadDir, compressedDir)
	imageHandler := handlers.NewImageHandler(imageService, uploadDir, compressedDir)

	// 基本路由
	r.GET("/", appHandler.HomePage)
	r.GET("/health", appHandler.HealthCheck)

	// 为前端兼容性添加直接路由
	r.POST("/upload", imageHandler.UploadAndCompress)   // 兼容原有前端调用
	r.GET("/images", imageHandler.ListCompressedImages) // 兼容原有前端调用

	// 提供静态文件访问
	r.Static("/static", uploadDir)
	r.Static("/compressed", compressedDir)

	// API 路由组
	api := r.Group("/api")
	{
		// 图片上传路由 - 兼容前端调用
		api.POST("/upload", imageHandler.UploadImage)                // 纯上传功能
		api.POST("/compress", imageHandler.CompressImage)            // 纯压缩功能
		api.POST("/upload-compress", imageHandler.UploadAndCompress) // 上传并压缩

		// 为前端兼容性提供静态文件访问
		api.Static("/static", compressedDir) // 前端期望通过 /api/static/ 访问压缩后的图片
		api.Static("/uploads", uploadDir)    // 访问原始上传文件
	}

	// API v1 路由组
	apiV1 := r.Group("/api/v1")
	{
		// 用户相关路由
		users := apiV1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.POST("", userHandler.CreateUser)
		}

		// 图片相关路由
		images := apiV1.Group("/images")
		{
			images.POST("/compress", imageHandler.UploadAndCompress)           // 上传并压缩图片
			images.GET("/formats", imageHandler.GetSupportedFormats)           // 获取支持的格式
			images.GET("/list", imageHandler.ListCompressedImages)             // 列出所有压缩图片
			images.GET("/download/:filename", imageHandler.DownloadCompressed) // 下载压缩图片
			images.DELETE("/:filename", imageHandler.DeleteCompressedImage)    // 删除压缩图片
		}
	}

	return r
}

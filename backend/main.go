package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	// 创建上传目录
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		log.Fatalf("创建上传目录失败: %v", err)
	}

	r := gin.Default()

	// 设置文件上传大小限制 (默认32MB，这里设置为10MB)
	r.MaxMultipartMemory = 10 << 20 // 10 MB

	// 提供静态文件访问
	r.Static("/static", "./uploads")

	// 文件上传路由
	r.POST("/upload", func(c *gin.Context) {
		// 获取上传的文件
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("获取文件失败: %s", err.Error()),
			})
			return
		}

		// 验证文件类型
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/jpg":  true,
			"image/png":  true,
			"image/gif":  true,
			"image/webp": true,
		}

		if !allowedTypes[file.Header.Get("Content-Type")] {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "不支持的文件类型，仅支持 JPEG, PNG, GIF, WEBP",
			})
			return
		}

		// 验证文件大小 (5MB)
		if file.Size > 5<<20 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "文件大小不能超过 5MB",
			})
			return
		}

		// 生成唯一文件名，防止冲突
		fileExt := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)
		dst := filepath.Join("./uploads", newFileName)

		// 保存文件
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": fmt.Sprintf("保存文件失败: %s", err.Error()),
			})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "文件上传成功",
			"data": gin.H{
				"fileName": file.Filename,
				"filePath": newFileName,
				"fileSize": file.Size,
				"fileType": file.Header.Get("Content-Type"),
			},
		})
	})

	// 获取上传文件列表
	r.GET("/images", func(c *gin.Context) {
		var files []gin.H

		err := filepath.Walk("./uploads", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasPrefix(info.Mode().String(), "-") {
				files = append(files, gin.H{
					"name": info.Name(),
					"size": info.Size(),
					"url":  fmt.Sprintf("/static/%s", info.Name()),
					"time": info.ModTime().Format(time.RFC3339),
				})
			}
			return nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": fmt.Sprintf("读取文件列表失败: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    files,
		})
	})

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "文件上传服务运行正常",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// 启动服务器
	fmt.Println("服务器启动在 http://localhost:8080")
	fmt.Println("上传端点: POST http://localhost:8080/upload")
	fmt.Println("静态文件访问: http://localhost:8080/static/文件名")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

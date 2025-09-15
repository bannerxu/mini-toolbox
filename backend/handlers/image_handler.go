package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"mini-toolbox/models"
	"mini-toolbox/utils"

	"github.com/gin-gonic/gin"
)

// ImageHandler 图片处理器
type ImageHandler struct {
	imageService  models.ImageService
	uploadDir     string
	compressedDir string
	maxFileSize   int64 // 最大文件大小（字节）
}

// NewImageHandler 创建新的图片处理器
func NewImageHandler(imageService models.ImageService, uploadDir, compressedDir string) *ImageHandler {
	return &ImageHandler{
		imageService:  imageService,
		uploadDir:     uploadDir,
		compressedDir: compressedDir,
		maxFileSize:   10 * 1024 * 1024, // 10MB
	}
}

// UploadImage 仅上传图片，不进行压缩
func (h *ImageHandler) UploadImage(c *gin.Context) {
	// 解析表单数据
	err := c.Request.ParseMultipartForm(h.maxFileSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: "文件太大或请求格式错误",
		})
		return
	}

	// 获取上传的文件
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: "请选择要上传的图片文件",
		})
		return
	}
	defer file.Close()

	// 验证文件格式
	if !h.imageService.ValidateImageFormat(fileHeader.Filename) {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: fmt.Sprintf("不支持的文件格式，支持的格式: %v", h.imageService.GetSupportedFormats()),
		})
		return
	}

	// 检查文件大小
	if fileHeader.Size > h.maxFileSize {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: fmt.Sprintf("文件大小超过限制 %d MB", h.maxFileSize/(1024*1024)),
		})
		return
	}

	// 生成唯一文件名
	timestamp := time.Now().Unix()
	originalFilename := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)
	inputPath := filepath.Join(h.uploadDir, originalFilename)

	// 保存上传的文件
	err = c.SaveUploadedFile(fileHeader, inputPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.LegacyErrorResponse{
			Success: false,
			Message: "保存文件失败",
		})
		return
	}

	// 返回上传成功信息
	c.JSON(http.StatusOK, utils.LegacySuccessResponse{
		Success: true,
		Message: "图片上传成功",
		Data: gin.H{
			"fileName":     fileHeader.Filename,
			"filePath":     originalFilename,
			"fileSize":     fileHeader.Size,
			"fileType":     fileHeader.Header.Get("Content-Type"),
			"originalName": fileHeader.Filename,
			"uploadTime":   time.Now().Unix(),
		},
	})
}

// CompressImage 压缩已上传的图片
func (h *ImageHandler) CompressImage(c *gin.Context) {
	// 获取要压缩的文件名
	filename := c.PostForm("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: "请提供要压缩的文件名",
		})
		return
	}

	// 构建文件路径
	inputPath := filepath.Join(h.uploadDir, filename)

	// 检查文件是否存在
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, utils.LegacyErrorResponse{
			Success: false,
			Message: "要压缩的文件不存在",
		})
		return
	}

	// 验证文件格式
	if !h.imageService.ValidateImageFormat(filename) {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: fmt.Sprintf("不支持的文件格式，支持的格式: %v", h.imageService.GetSupportedFormats()),
		})
		return
	}

	// 获取压缩选项
	options := h.parseCompressionOptions(c)

	// 生成压缩后文件名
	compressedFilename := models.GenerateUniqueFilename(filename)
	outputPath := filepath.Join(h.compressedDir, compressedFilename)

	// 压缩图片
	result, err := h.imageService.CompressImage(inputPath, outputPath, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.LegacyErrorResponse{
			Success: false,
			Message: fmt.Sprintf("图片压缩失败: %v", err),
		})
		return
	}

	// 返回压缩结果
	c.JSON(http.StatusOK, utils.LegacySuccessResponse{
		Success: true,
		Message: "图片压缩成功",
		Data: gin.H{
			"originalFile":     filename,
			"compressedFile":   compressedFilename,
			"originalSize":     result.OriginalSize,
			"compressedSize":   result.CompressedSize,
			"compressionRatio": result.Ratio,
			"quality":          options.Quality,
			"width":            options.Width,
			"height":           options.Height,
			"originalUrl":      fmt.Sprintf("/api/uploads/%s", filename),
			"compressedUrl":    fmt.Sprintf("/api/static/%s", compressedFilename),
		},
	})
}

// UploadAndCompress 上传并压缩图片
func (h *ImageHandler) UploadAndCompress(c *gin.Context) {
	// 解析表单数据
	err := c.Request.ParseMultipartForm(h.maxFileSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: "文件太大或请求格式错误",
		})
		return
	}

	// 获取上传的文件
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: "请选择要上传的图片文件",
		})
		return
	}
	defer file.Close()

	// 验证文件格式
	if !h.imageService.ValidateImageFormat(fileHeader.Filename) {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: fmt.Sprintf("不支持的文件格式，支持的格式: %v", h.imageService.GetSupportedFormats()),
		})
		return
	}

	// 检查文件大小
	if fileHeader.Size > h.maxFileSize {
		c.JSON(http.StatusBadRequest, utils.LegacyErrorResponse{
			Success: false,
			Message: fmt.Sprintf("文件大小超过限制 %d MB", h.maxFileSize/(1024*1024)),
		})
		return
	}

	// 生成唯一文件名
	timestamp := time.Now().Unix()
	originalFilename := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)
	inputPath := filepath.Join(h.uploadDir, originalFilename)

	// 保存上传的文件
	err = c.SaveUploadedFile(fileHeader, inputPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.LegacyErrorResponse{
			Success: false,
			Message: "保存文件失败",
		})
		return
	}

	// 获取压缩选项
	options := h.parseCompressionOptions(c)

	// 生成压缩后文件名
	compressedFilename := models.GenerateUniqueFilename(originalFilename)
	outputPath := filepath.Join(h.compressedDir, compressedFilename)

	// 压缩图片
	result, err := h.imageService.CompressImage(inputPath, outputPath, options)
	if err != nil {
		// 清理上传的文件
		os.Remove(inputPath)
		c.JSON(http.StatusInternalServerError, utils.LegacyErrorResponse{
			Success: false,
			Message: fmt.Sprintf("图片压缩失败: %v", err),
		})
		return
	}

	// 清理原始上传文件（可选，这里保留以供对比）
	// os.Remove(inputPath)

	// 为前端兼容性，返回期望的格式
	c.JSON(http.StatusOK, utils.LegacySuccessResponse{
		Success: true,
		Message: "图片上传成功",
		Data: gin.H{
			"fileName":         fileHeader.Filename,
			"filePath":         compressedFilename,
			"fileSize":         result.CompressedSize,
			"fileType":         fileHeader.Header.Get("Content-Type"),
			"originalSize":     result.OriginalSize,
			"compressionRatio": result.Ratio,
		},
	})
}

// GetSupportedFormats 获取支持的图片格式
func (h *ImageHandler) GetSupportedFormats(c *gin.Context) {
	formats := h.imageService.GetSupportedFormats()
	c.JSON(http.StatusOK, gin.H{
		"supportedFormats": formats,
		"maxFileSize":      fmt.Sprintf("%d MB", h.maxFileSize/(1024*1024)),
	})
}

// DownloadCompressed 下载压缩后的图片
func (h *ImageHandler) DownloadCompressed(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ResponseError{
			Error: "缺少文件名",
		})
		return
	}

	filePath := filepath.Join(h.compressedDir, filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, utils.ResponseError{
			Error: "文件不存在",
		})
		return
	}

	// 设置响应头
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "application/octet-stream")

	// 发送文件
	c.File(filePath)
}

// ListCompressedImages 列出所有压缩的图片
func (h *ImageHandler) ListCompressedImages(c *gin.Context) {
	files, err := os.ReadDir(h.compressedDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Error: "读取压缩文件目录失败",
		})
		return
	}

	var images []map[string]interface{}
	for _, file := range files {
		if !file.IsDir() && h.imageService.ValidateImageFormat(file.Name()) {
			info, _ := file.Info()
			images = append(images, map[string]interface{}{
				"filename":    file.Name(),
				"size":        info.Size(),
				"modTime":     info.ModTime(),
				"downloadUrl": fmt.Sprintf("/api/v1/images/download/%s", file.Name()),
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"images": images,
		"count":  len(images),
	})
}

// DeleteCompressedImage 删除压缩的图片
func (h *ImageHandler) DeleteCompressedImage(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ResponseError{
			Error: "缺少文件名",
		})
		return
	}

	filePath := filepath.Join(h.compressedDir, filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, utils.ResponseError{
			Error: "文件不存在",
		})
		return
	}

	// 删除文件
	err := os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Error: "删除文件失败",
		})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess{
		Message: "文件删除成功",
		Data:    gin.H{"filename": filename},
	})
}

// parseCompressionOptions 解析压缩选项
func (h *ImageHandler) parseCompressionOptions(c *gin.Context) models.CompressionOption {
	options := models.CompressionOption{
		Quality:    85,   // 默认质量
		Width:      0,    // 默认不调整宽度
		Height:     0,    // 默认不调整高度
		KeepAspect: true, // 默认保持宽高比
	}

	// 解析质量参数
	if qualityStr := c.PostForm("quality"); qualityStr != "" {
		if quality, err := strconv.Atoi(qualityStr); err == nil && quality > 0 && quality <= 100 {
			options.Quality = quality
		}
	}

	// 解析宽度参数
	if widthStr := c.PostForm("width"); widthStr != "" {
		if width, err := strconv.ParseUint(widthStr, 10, 32); err == nil {
			options.Width = uint(width)
		}
	}

	// 解析高度参数
	if heightStr := c.PostForm("height"); heightStr != "" {
		if height, err := strconv.ParseUint(heightStr, 10, 32); err == nil {
			options.Height = uint(height)
		}
	}

	// 解析是否保持宽高比
	if keepAspectStr := c.PostForm("keepAspect"); keepAspectStr != "" {
		if keepAspect, err := strconv.ParseBool(keepAspectStr); err == nil {
			options.KeepAspect = keepAspect
		}
	}

	return options
}

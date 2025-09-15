package models

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

// CompressionOption 压缩选项
type CompressionOption struct {
	Quality    int  `json:"quality"`    // JPEG 质量 (1-100)
	Width      uint `json:"width"`      // 目标宽度，0 表示保持原比例
	Height     uint `json:"height"`     // 目标高度，0 表示保持原比例
	KeepAspect bool `json:"keepAspect"` // 是否保持宽高比
}

// CompressResult 压缩结果
type CompressResult struct {
	OriginalSize   int64  `json:"originalSize"`   // 原始文件大小（字节）
	CompressedSize int64  `json:"compressedSize"` // 压缩后文件大小（字节）
	Filename       string `json:"filename"`       // 压缩后文件名
	CompressionURL string `json:"compressionUrl"` // 压缩后文件访问URL
	Ratio          string `json:"ratio"`          // 压缩比例
}

// ImageService 图片服务接口
type ImageService interface {
	CompressImage(inputPath, outputPath string, options CompressionOption) (*CompressResult, error)
	GetSupportedFormats() []string
	ValidateImageFormat(filename string) bool
}

// DefaultImageService 默认图片服务实现
type DefaultImageService struct {
	supportedFormats []string
	uploadDir        string
	compressedDir    string
}

// NewDefaultImageService 创建默认图片服务
func NewDefaultImageService(uploadDir, compressedDir string) *DefaultImageService {
	return &DefaultImageService{
		supportedFormats: []string{".jpg", ".jpeg", ".png"},
		uploadDir:        uploadDir,
		compressedDir:    compressedDir,
	}
}

// GetSupportedFormats 获取支持的图片格式
func (s *DefaultImageService) GetSupportedFormats() []string {
	return s.supportedFormats
}

// ValidateImageFormat 验证图片格式
func (s *DefaultImageService) ValidateImageFormat(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, format := range s.supportedFormats {
		if ext == format {
			return true
		}
	}
	return false
}

// CompressImage 压缩图片
func (s *DefaultImageService) CompressImage(inputPath, outputPath string, options CompressionOption) (*CompressResult, error) {
	// 获取原始文件信息
	originalInfo, err := os.Stat(inputPath)
	if err != nil {
		return nil, fmt.Errorf("无法获取原始文件信息: %v", err)
	}

	// 打开原始图片
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("无法打开输入文件: %v", err)
	}
	defer inputFile.Close()

	// 解码图片
	img, format, err := image.Decode(inputFile)
	if err != nil {
		return nil, fmt.Errorf("无法解码图片: %v", err)
	}

	// 调整图片尺寸
	if options.Width > 0 || options.Height > 0 {
		if options.KeepAspect {
			// 保持宽高比，使用 imaging 库
			if options.Width > 0 && options.Height > 0 {
				img = imaging.Fit(img, int(options.Width), int(options.Height), imaging.Lanczos)
			} else if options.Width > 0 {
				img = imaging.Resize(img, int(options.Width), 0, imaging.Lanczos)
			} else if options.Height > 0 {
				img = imaging.Resize(img, 0, int(options.Height), imaging.Lanczos)
			}
		} else {
			// 不保持宽高比，直接调整尺寸
			img = resize.Resize(options.Width, options.Height, img, resize.Lanczos3)
		}
	}

	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return nil, fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outputFile.Close()

	// 根据格式编码图片
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		quality := options.Quality
		if quality <= 0 || quality > 100 {
			quality = 85 // 默认质量
		}
		err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: quality})
	case "png":
		err = png.Encode(outputFile, img)
	default:
		return nil, errors.New("不支持的图片格式")
	}

	if err != nil {
		return nil, fmt.Errorf("编码图片失败: %v", err)
	}

	// 获取压缩后文件信息
	compressedInfo, err := os.Stat(outputPath)
	if err != nil {
		return nil, fmt.Errorf("无法获取压缩文件信息: %v", err)
	}

	// 计算压缩比例
	compressionRatio := float64(compressedInfo.Size()) / float64(originalInfo.Size()) * 100

	result := &CompressResult{
		OriginalSize:   originalInfo.Size(),
		CompressedSize: compressedInfo.Size(),
		Filename:       filepath.Base(outputPath),
		CompressionURL: fmt.Sprintf("/api/v1/images/download/%s", filepath.Base(outputPath)),
		Ratio:          fmt.Sprintf("%.1f%%", compressionRatio),
	}

	return result, nil
}

// GenerateUniqueFilename 生成唯一文件名
func GenerateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	name := strings.TrimSuffix(originalFilename, ext)
	return fmt.Sprintf("%s_compressed_%d%s", name, generateTimestamp(), ext)
}

// generateTimestamp 生成时间戳
func generateTimestamp() int64 {
	return time.Now().Unix()
}

// CopyFile 复制文件
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

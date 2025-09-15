package utils

import (
	"strconv"
)

// ParseID 解析字符串ID为整数
func ParseID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ResponseError 错误响应结构
type ResponseError struct {
	Error string `json:"error"`
}

// ResponseSuccess 成功响应结构
type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// LegacySuccessResponse 兼容前端的成功响应结构
type LegacySuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// LegacyErrorResponse 兼容前端的错误响应结构
type LegacyErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

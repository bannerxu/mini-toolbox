package models

import "errors"

// 定义应用程序中的错误
var (
	ErrUserNotFound = errors.New("用户未找到")
	ErrInvalidInput = errors.New("无效的输入数据")
)

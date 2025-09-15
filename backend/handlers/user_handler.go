package handlers

import (
	"net/http"

	"mini-toolbox/models"
	"mini-toolbox/utils"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService models.UserService
}

// NewUserHandler 创建新的用户处理器
func NewUserHandler(userService models.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers 获取所有用户
func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	response := models.UsersResponse{
		Users: users,
		Count: len(users),
	}
	c.JSON(http.StatusOK, response)
}

// GetUserByID 根据ID获取用户
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ParseID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{
			Error: "无效的用户ID",
		})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		if err == models.ErrUserNotFound {
			c.JSON(http.StatusNotFound, utils.ResponseError{
				Error: "用户未找到",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Error: "服务器内部错误",
		})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{User: *user})
}

// CreateUser 创建新用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError{
			Error: "无效的请求数据",
		})
		return
	}

	user := h.userService.CreateUser(req.Name, req.Age)

	c.JSON(http.StatusCreated, utils.ResponseSuccess{
		Message: "用户创建成功",
		Data:    user,
	})
}

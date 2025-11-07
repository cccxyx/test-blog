package handlers

import (
	"github.com/dingdinglz/test-blog/models"
	"github.com/dingdinglz/test-blog/services"
	"github.com/dingdinglz/test-blog/utils"
	"github.com/gin-gonic/gin"
)

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误: "+err.Error())
		return
	}

	// 调用服务层
	user, err := services.Register(req.Username, req.Password, req.Email)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	// 返回响应
	utils.Success(c, gin.H{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	}, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "参数错误: "+err.Error())
		return
	}

	// 调用服务层
	token, user, err := services.Login(req.Username, req.Password)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	// 返回响应
	utils.Success(c, gin.H{
		"token": token,
		"user": models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, "登录成功")
}

// GetInfo 获取当前用户信息
func GetInfo(c *gin.Context) {
	// 从Context获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, "未授权")
		return
	}

	// 调用服务层
	user, err := services.GetUserByID(userID.(uint))
	if err != nil {
		utils.Error(c, 404, err.Error())
		return
	}

	// 返回响应
	utils.Success(c, models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, "success")
}

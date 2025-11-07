package services

import (
	"errors"

	"github.com/dingdinglz/test-blog/database"
	"github.com/dingdinglz/test-blog/models"
	"github.com/dingdinglz/test-blog/utils"
	"gorm.io/gorm"
)

// Register 用户注册
func Register(username, password, email string) (*models.User, error) {
	db := database.GetDB()

	// 检查用户名是否已存在
	var existUser models.User
	if err := db.Where("username = ?", username).First(&existUser).Error; err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	if err := db.Where("email = ?", email).First(&existUser).Error; err == nil {
		return nil, errors.New("邮箱已被注册")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户
	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	if err := db.Create(user).Error; err != nil {
		return nil, errors.New("用户创建失败")
	}

	return user, nil
}

// Login 用户登录
func Login(username, password string) (string, *models.User, error) {
	db := database.GetDB()

	// 查找用户
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil, errors.New("用户不存在")
		}
		return "", nil, errors.New("查询用户失败")
	}

	// 验证密码
	if !utils.CheckPassword(user.Password, password) {
		return "", nil, errors.New("密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", nil, errors.New("生成token失败")
	}

	return token, &user, nil
}

// GetUserByID 根据ID获取用户信息
func GetUserByID(userID uint) (*models.User, error) {
	db := database.GetDB()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, errors.New("查询用户失败")
	}

	return &user, nil
}

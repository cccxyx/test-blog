package models

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"`
	Email    string    `gorm:"uniqueIndex;not null" json:"email"`
	Articles []Article `gorm:"foreignKey:UserID" json:"articles,omitempty"`
}

// UserResponse 用户响应结构（不包含密码）
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

package models

import (
	"errors"
)

// user model本体
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required,gt=2,lt=10"`
	Password string `json:"password" binding:"required"`
}

// 定义
var (
	errNameEmpty = errors.New("name is empty")
)

// 添加用户api接收的model
type UserAddInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 更新用户api接收的model
type UserUpdateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

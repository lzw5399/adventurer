package model

// user model本体
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required,gt=2,lt=10"`
	Password string `json:"password" binding:"required"`
}

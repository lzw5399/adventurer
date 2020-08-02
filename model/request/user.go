/**
 * @Author: lzw5399
 * @Date: 2020/8/2 13:59
 * @Desc: user相关api的request model
 */
package request

// 添加用户api接收的model
type UserAddRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 更新用户api接收的model
type UserUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

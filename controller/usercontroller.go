package controller

import (
	"adventurer/global/response"
	"adventurer/model"
	req "adventurer/model/request"
	"adventurer/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

// @Tags account
// @Summary 用户注册
// @Accept  json
// @Produce json
// @Param   user body request.UserRegisterRequest true "register"
// @Success 200 {object} response.HTTPResponse
// @Router /account/register [post]
func Register(c *gin.Context) {
	var r req.UserRegisterRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage(c, http.StatusBadRequest, err)
		return
	}
}

// @Tags account
// @Summary 用户登陆
// @Accept  json
// @Produce json
// @Param   user body request.UserLoginRequest true "username&pwd"
// @Success 200 {string} string	"token"
// @Router /account/login [post]
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage(c, http.StatusBadRequest, err)
		return
	}

	if token, success := service.Login(user); success {
		c.String(http.StatusOK, token)
	} else {
		response.FailWithMessage(c, http.StatusBadRequest, "username or password invalid")
	}
}

// @Tags account
// @Summary 获取当前的用户信息
// @Accept  json
// @Produce json
// @Success 200 {object} model.CustomClaims
// @Security ApiKeyAuth
// @Router /account/user [get]
func GetCurrentUser(c *gin.Context) {
	a, _ := c.Get("currentUser")

	if obj, success := a.(*model.CustomClaims); !success {
		response.FailWithMessage(c, http.StatusNotFound, "user not found")
	} else {
		response.OkWithData(c, obj)
	}
}

func ListUsers(c *gin.Context) {

}

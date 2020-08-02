/**
 * @Author: lzw5399
 * @Date: 2020/7/31 23:21
 * @Desc: 格式化错误响应
 */
package response

import (
	"adventurer/global"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	// 全局翻译器
	trans ut.Translator
)

type HTTPResponse struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func init() {
	_ = initTrans(global.DL_CONFIG.System.Locate)
}

// initTrans 初始化翻译器
func initTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func Ok(c *gin.Context) {
	result(c, true, nil, "success")
}

func OkWithMessage(message string, c *gin.Context) {
	result(c, true, nil, message)
}

func OkWithData(data interface{}, c *gin.Context) {
	result(c, true, data, "操作成功")
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	result(c, true, data, message)
}

func Failed(c *gin.Context, status int) {
	FailWithMessage(c, status, "操作失败")
}

func FailWithMessage(c *gin.Context, status int, err interface{}) {
	var msg interface{}

	switch err.(type) {
	case validator.ValidationErrors:
		msg = err.(validator.ValidationErrors).Translate(trans)

	case error:
		msg = err.(error).Error()

	case string:
		msg = err.(string)

	default:
		msg = ""
	}

	// TODO: 这一块之后需要根据status code来做判断
	// 500: 的话不能直接返回，而是记日志再返回通用的错误信息
	switch status {
	// 400
	case http.StatusBadRequest:

	// 401
	case http.StatusUnauthorized:

	// 404
	case http.StatusNotFound:

	// 500
	case http.StatusInternalServerError:
		msg = "server internal error, please contact the maintainer"
	}

	resultWithStatus(c, status, false, nil, msg)
	c.Abort()
}

func result(c *gin.Context, success bool, data interface{}, msg string) {
	c.JSON(http.StatusOK, HTTPResponse{
		success,
		data,
		msg,
	})
}

func resultWithStatus(c *gin.Context, statusCode int, success bool, data interface{}, msg interface{}) {
	c.JSON(statusCode, HTTPResponse{
		success,
		msg,
		data,
	})
}

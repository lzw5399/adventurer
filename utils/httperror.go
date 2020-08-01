/**
 * @Author: lzw5399
 * @Date: 2020/7/31 23:21
 * @Desc: 格式化错误响应
 */
package utils

import (
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

type HTTPError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func init() {
	_ = initTrans(Config.Locate)
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

func NewError(ctx *gin.Context, status int, err interface{}) {
	er := HTTPError{
		Code:    status,
	}

	switch err.(type) {
	case validator.ValidationErrors:
		er.Message = err.(validator.ValidationErrors).Translate(trans)

	case error:
		er.Message = err.(error).Error()

	case string:
		er.Message = err.(string)

	default:
		er.Message = ""
	}

	// TODO: 这一块之后需要根据status code来做判断
	// 400: 可以不用做什么处理，前面已经处理过了
	// 401：
	// 404：
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
		er.Message = "server internal error, please contact the maintainer"
	}

	ctx.JSON(status, er)
	ctx.Abort()
}

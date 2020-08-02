/**
 * @Author: lzw5399
 * @Date: 2020/7/30 13:30
 * @Desc: load settings when application starts
 */
package initilize

import (
	"fmt"
	"os"
	"strings"

	"dl-admin-go/global"
	"dl-admin-go/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

// go同一个包内有多个init，则会按文件名顺序执行
// 如果其他的init会依赖当前这个config的init（最起码db.go依赖）
// 所以如果需要依赖，不要让package内的其他文件排在config.go前面
func init() {
	envcode := getEnvCode()
	overrideConfigFileName := fmt.Sprintf("config/appsettings.%s.yaml", envcode)

	var err error
	if util.Exists(overrideConfigFileName) {
		err = configor.Load(&global.DL_CONFIG, "config/appsettings.yaml", overrideConfigFileName)
	} else {
		err = configor.Load(&global.DL_CONFIG, "config/appsettings.yaml")
	}

	if err != nil {
		panic("resolve settings failed...")
	}
}

var envMap = map[string]string{
	"debug":   "Development",
	"release": "Production",
}

func getEnvCode() string {
	ginmode := os.Getenv(gin.EnvGinMode)

	for k, v := range envMap {
		if k == strings.ToLower(ginmode) {
			return v
		}
	}

	return "Development"
}

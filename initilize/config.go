/**
 * @Author: lzw5399
 * @Date: 2020/7/30 13:30
 * @Desc: load settings when application starts
 */
package initilize

import (
	"dl-admin-go/config"
	"fmt"
	"os"
	"strings"

	"dl-admin-go/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

var Config config.Config

func init() {
	envcode := getEnvCode()
	overrideConfigFileName := fmt.Sprintf("config/appsettings.%s.yaml", envcode)

	var err error
	if util.Exists(overrideConfigFileName) {
		err = configor.Load(&Config, "config/appsettings.yaml", overrideConfigFileName)
	} else {
		err = configor.Load(&Config, "config/appsettings.yaml")
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

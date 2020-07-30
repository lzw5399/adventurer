/**
 * @Author: lzw5399
 * @Date: 2020/7/30 13:30
 * @Desc: load settings when application starts
 */
package utils

import (
	"dl-admin-go/models/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"os"
	"strings"
)

var Config configs.Config

func init(){
	loadSettings()
}

func loadSettings() {
	envcode := getEnvCode()
	overrideConfigFileName := fmt.Sprintf("configs/appsettings.%s.yaml", envcode)

	var err error
	if Exists(overrideConfigFileName) {
		err = configor.Load(&Config, "configs/appsettings.yaml", overrideConfigFileName)
	} else {
		err = configor.Load(&Config, "configs/appsettings.yaml")
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

/**
 * @Author: lzw5399
 * @Date: 2020/8/2 13:19
 * @Desc: 用于存放全局对象，且以下对象会在应用初始化时也被填充
 */
package global

import (
	"adventurer/config"
	"github.com/jinzhu/gorm"
)

var (
	DL_CONFIG config.Config
	DL_DB     *gorm.DB
)

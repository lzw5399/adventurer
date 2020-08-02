/**
 * @Author: lzw5399
 * @Date: 2020/8/2 14:55
 * @Desc: 初始化数据库并产生数据库全局变量, 这里默认是postgres
 */
package initilize

import (
	"dl-admin-go/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	dbconfig := global.DL_CONFIG.Db

	connstr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.Username, dbconfig.InitialDb, dbconfig.Password)
	if conn, err := gorm.Open("postgres", connstr); err != nil {
		// todo 这里要记录log
		panic(err)
	} else {
		global.DL_DB = conn
		global.DL_DB.DB().SetMaxIdleConns(dbconfig.MaxIdleConn)
		global.DL_DB.DB().SetMaxOpenConns(dbconfig.MaxOpenConn)
		global.DL_DB.LogMode(dbconfig.LogMode)
	}
}

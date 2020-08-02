package main

import (
	"adventurer/global"
	"fmt"
	"os"

	_ "adventurer/docs"
	_ "adventurer/initilize"
	"adventurer/router"
)

// @title Swagger Example API
// @version 1.0
// @description dl-admin backend api document

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := router.InitRouter()

	defer global.DL_DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("server listening at port: " + port)
	_ = r.Run(":" + port)
}

package main

import (
	"dl-admin-go/global"
	"fmt"
	"os"

	_ "dl-admin-go/docs"
	_ "dl-admin-go/initilize"
	"dl-admin-go/router"
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

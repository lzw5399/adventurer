package main

import (
	_ "dl-admin-go/docs"
	"dl-admin-go/routers"
	_ "dl-admin-go/utils"

	"fmt"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description dl-admin backend api document

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := routers.InitRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("server listening at port: " + port)
	_ = r.Run(":" + port)
}

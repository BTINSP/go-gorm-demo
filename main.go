package main

import (
	"gin-gorm-demo/database"
	"gin-gorm-demo/router"
)




func main() {

	//	初始化database
	database.Init()

	//	初始化router
	server := router.Init()

	//	启动server
	server.Run(":8080")
}



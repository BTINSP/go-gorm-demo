package main

import (
	"gin-grom-demo/database"
	"gin-grom-demo/router"
)




func main() {

	//	初始化database
	database.Init()

	//	初始化router
	server := router.Init()

	//	启动server
	server.Run(":8080")
}



package router

import (
	"gin-gorm-demo/handle"
	"gin-gorm-demo/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", Index)	

	//	user router
	userRouter := engine.Group("/user")
		userRouter.Use(middleware.Authorization)
		userRouter.GET("/get", handle.GetUserByIdHandle)
		userRouter.GET("/all", handle.GetAllUser)


	//	404 handling
	engine.NoRoute(NoRoute)
	return engine
}

func Index(ctx *gin.Context) {
	ctx.String(200, "Hello World.")
}

func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, "Not Found.")
}


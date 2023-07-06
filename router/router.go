package router

import (
	"gin-grom-demo/handle"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine {
	engine := gin.Default()

	engine.GET("/", Index)

	
	userRouter := engine.Group("/user")
	userRouter.GET("/get", handle.UserGetUserByIdHandle)


	engine.NoRoute(NoRoute)
	return engine
}

func Index(ctx *gin.Context) {
	ctx.String(200, "Hello World.")
}

func NoRoute (ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, "Not Found.")
}


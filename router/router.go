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

	userRouter := engine.Group("/user", middleware.Authorization)
	userRouter.GET("/get",handle.UserGetUserByIdHandle)


	engine.NoRoute(NoRoute)
	return engine
}

func Index(ctx *gin.Context) {
	ctx.String(200, "Hello World.")
}

func NoRoute (ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, "Not Found.")
}


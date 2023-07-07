package middleware

import (
	"github.com/gin-gonic/gin"
	result "gin-gorm-demo/common"
)



func Authorization(ctx *gin.Context) {
	auth := ctx.GetHeader("auth")
	if auth == "jwt" {
		ctx.Next()
		return
	}
	result.Fail(ctx, result.UnAuthorization, nil)
	ctx.Abort()
} 
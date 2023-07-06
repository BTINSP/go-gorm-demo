package middleware

import (
	"github.com/gin-gonic/gin"
	result "gin-gorm-demo/common"
)



func Authorization(ctx *gin.Context) {
	auth := ctx.GetHeader("auth")
	if auth == "jwt" {
		return
	}
	
	result.Fail(ctx, result.BadRequest, nil)
	return
} 
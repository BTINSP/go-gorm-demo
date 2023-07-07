package common

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	OK					= statusCode(200, "success")
	BadRequest			= statusCode(400, "fail")
	UnAuthorization		= statusCode(401, "请先登陆.")
)


type Response struct {
	Code int			`json:"code"`
	Message string		`json:"message"`
	Data interface{}	`json:"data"`
}

func statusCode(code int, message string) Response {
	return Response {
		code,
		message,
		nil,
	}
}


func Default(ctx *gin.Context, response Response, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		Response {
			response.Code,
			response.Message,
			data,
		},
	)
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		Response {
			200,
			"success",
			data,
		},
	)
}

func Fail(ctx *gin.Context, response Response, data interface{}) {
	ctx.JSON(
		http.StatusBadRequest,
		Response {
			response.Code,
			response.Message,
			data,
		},
	)
}
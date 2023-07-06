package handle

import (
	"gin-grom-demo/database"
	"github.com/gin-gonic/gin"
	modal "gin-grom-demo/model"
	result "gin-grom-demo/common"

)



//	Method: GET Path: /user/get
func UserGetUserByIdHandle(ctx *gin.Context) {

	id := ctx.Query("id")

	var user modal.User
	db := database.GetDB()
	db.Table("user").Find(&user, "id = ?", id)

	result.Default(ctx, result.OK, user)
}
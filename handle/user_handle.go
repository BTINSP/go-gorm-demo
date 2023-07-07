package handle

import (
	"gin-gorm-demo/database"
	"github.com/gin-gonic/gin"
	modal "gin-gorm-demo/model"
	result "gin-gorm-demo/common"

)



//	Method: GET Path: /user/get
func GetUserByIdHandle(ctx *gin.Context) {

	id := ctx.Query("id")

	var user modal.User
	db := database.GetDB()
	db.Table("user").Find(&user, "id = ?", id)

	result.Default(ctx, result.OK, user)
}

//	Method: GET Path: /user/all
func GetAllUser(ctx *gin.Context) {

	var users []modal.User
	database.GetDB().Table("user").Find(&users)

	result.Success(ctx, users)
}
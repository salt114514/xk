package userController

import (
	"STU/app/models"
	"STU/app/services/userService"
	"STU/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseError(c, 200501, "参数错误")
		return
	}

	err = userService.CheckUserExistByAccount(data.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseInternalError(c)
		} else {
			utils.ResponseError(c, 401, "账号不存在")
		}
		return
	}

	var user *models.User
	user, err = userService.GetUserByAccount(data.Account)
	if err != nil {
		utils.ResponseInternalError(c)
		return
	}

	flag := userService.ComparePwd(data.Password, user.Password)
	if !flag {
		utils.ResponseError(c, 402, "密码错误")
		return
	}

	utils.ResponseSuccess(c, user)
}

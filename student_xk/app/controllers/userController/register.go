package userController

import (
	"STU/app/models"
	"STU/app/services/userService"
	"STU/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Register struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Type     uint   `json:"type" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func RegistData(c *gin.Context) {
	var data Register
	err1 := c.ShouldBindJSON(&data)
	if err1 != nil {
		utils.ResponseError(c, 200501, "参数错误")
	}

	err1 = userService.CheckUserExistByAccount(data.Account)
	if err1 == nil {
		utils.ResponseError(c, 200504, "账号已存在")
		return
	} else if err1 != nil && err1 != gorm.ErrRecordNotFound {
		utils.ResponseInternalError(c)
		return
	}

	account := data.Account
	_, err2 := strconv.Atoi(account)
	if err2 != nil {
		utils.Response(c, 200, 400, "账号格式不正确", nil)
		return
	}

	if len(data.Password) < 8 || len(data.Password) > 20 {
		utils.Response(c, 200, 401, "密码长度必须在8~20位之间", nil)
		return
	}

	err := userService.Register(models.User{
		Account:  data.Account,
		Password: data.Password,
		Type:     data.Type,
		Name:     data.Name,
	})
	if err != nil {
		utils.ResponseInternalError(c)
		return
	}

	utils.ResponseSuccess(c, nil)

}

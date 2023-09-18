package teacherConcroller

import (
	"STU/app/models"
	"STU/app/services/teacherService"
	"STU/app/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 增
type CourseData struct {
	UserID    uint   `json:"user_id" bingding:"required"`
	ClassName string `json:"class_name" binding:"required"`
	Time      uint   `json:"time" binding:"required"`
	Weekday   uint   `json:"weekday" binding:"required"`
	Type      uint   `json:"type" binding:"required"`
	Number    uint   `json:"number" binding:"required"`
}

func CreateCourse(c *gin.Context) {
	var data CourseData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseError(c, 200501, "参数错误")
	}

	err = teacherService.CreateCourse(models.Course{
		UserID:    data.UserID,
		ClassName: data.ClassName,
		Time:      data.Time,
		Weekday:   data.Weekday,
		Type:      data.Type,
		Number:    data.Number,
	})

	if err != nil {
		utils.ResponseInternalError(c)
		return
	}

	utils.ResponseSuccess(c, nil)
}

// 改
type UpdateCourseData struct {
	UserID    uint   `json:"user_id" bingding:"required"`
	ClassID   uint   `json:"class_id" bingding:"required"`
	ClassName string `json:"class_name"`
	Time      uint   `json:"time"`
	Weekday   uint   `json:"weekday"`
	Type      uint   `json:"type"`
	Number    uint   `json:"number"`
}

func UpdateCourse(c *gin.Context) {
	var data UpdateCourseData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseError(c, 200501, "参数错误")
		return
	}

	err = teacherService.UpdateCourse(models.Course{
		UserID:    data.UserID,
		ClassID:   data.ClassID,
		ClassName: data.ClassName,
		Time:      data.Time,
		Weekday:   data.Weekday,
		Type:      data.Type,
		Number:    data.Number,
	})
	if err != nil {
		utils.ResponseInternalError(c)
		return
	}

	utils.ResponseSuccess(c, nil)
}

// 删
type DeleteCourseData struct {
	UserID  uint `json:"user_id" binding:"required"`
	ClassID uint `json:"class_id" binding:"required"`
}

func DeleteCourse(c *gin.Context) {
	var data DeleteCourseData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ResponseError(c, 200501, "参数错误")
		return
	}

	var user *models.Course
	user, err = teacherService.GetUserByUserID(data.ClassID)
	if err != nil {
		utils.ResponseInternalError(c)
		fmt.Println("1")
		return
	}

	flag := teacherService.CompareUser(data.UserID, user.UserID)
	if !flag {
		utils.ResponseError(c, 402, "无权限删除，不是同一个用户")
		return
	}

	err = teacherService.DeleteCourse(data.ClassID)
	if err != nil {
		utils.ResponseInternalError(c)
		return
	}

	utils.ResponseSuccess(c, nil)
}

// 查
type GetCourseData struct {
	UserID uint `form:"user_id" binding:"required"`
}

func GetCourse(c *gin.Context) {
	var data GetCourseData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.ResponseError(c, 200501, "参数错误")
		return
	}

	var courseList []models.Course
	courseList, err = teacherService.GetCourseUpdated(data.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseError(c, 200506, "未找到教师")
			return
		} else {
			utils.ResponseInternalError(c)
			return
		}
	}

	utils.ResponseSuccess(c, gin.H{
		"class_list": courseList,
	})
}

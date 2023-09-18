package router

import (
	"STU/app/controllers/teacherConcroller"
	"STU/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	user := api.Group("/user")
	{
		user.POST("/login", userController.Login)
		user.POST("/reg", userController.RegistData)
	}

	teacher := api.Group("/teacher")
	course := teacher.Group("/course")
	{
		course.POST("", teacherConcroller.CreateCourse)
		course.PUT("", teacherConcroller.UpdateCourse)
		course.DELETE("", teacherConcroller.DeleteCourse)
		course.GET("", teacherConcroller.GetCourse)
	}
}

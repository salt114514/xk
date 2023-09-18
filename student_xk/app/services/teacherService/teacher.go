package teacherService

import (
	"STU/app/models"
	"STU/config/database"
)

func CreateCourse(course models.Course) error {
	result := database.DB.Create(&course)
	return result.Error
}

func UpdateCourse(course models.Course) error {
	result := database.DB.Debug().Omit("user_id", "total").Save(&course)
	return result.Error
}

func DeleteCourse(id uint) error {
	result := database.DB.Debug().Where("class_id = ?", id).Delete(&models.Course{})
	return result.Error
}

func GetCourseUpdated(id uint) ([]models.Course, error) {
	result := database.DB.Debug().Where("user_id = ?", id).First(&models.Course{})
	if result.Error != nil {
		return nil, result.Error
	}
	var courseList []models.Course
	result = database.DB.Debug().Where("user_id = ?", id).Find(&courseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return courseList, nil
}

func GetUserByUserID(ClassID uint) (*models.Course, error) {
	var user models.Course
	result := database.DB.Where("class_id=?", ClassID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CompareUser(user1 uint, user2 uint) bool {
	return user1 == user2
}

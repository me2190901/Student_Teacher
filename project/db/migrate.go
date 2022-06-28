package db

import(
	"project/models"
)
func Migrate(){
	DbConn.AutoMigrate(&models.User{})
	DbConn.AutoMigrate(&models.Student{})
	DbConn.AutoMigrate(&models.Teacher{})
	DbConn.AutoMigrate(&models.Classroom{})
	DbConn.AutoMigrate(&models.UserMobile{})
	DbConn.AutoMigrate(&models.ClassroomTeachers{})
	DbConn.AutoMigrate(&models.ClassroomStudents{})
	models.Roles = map[string]bool{
		"admin": true, "manager": true, "teacher": true, "student": true,
	}
}
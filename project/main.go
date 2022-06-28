package main

import (
	"project/common"
	"project/controllers"
	"project/db"
	"github.com/gin-gonic/gin"
)

func main() {
	if err:=db.Connect();err!=nil{
		return
	}
	db.Migrate()
	r := gin.Default()
	r.POST("/callback/", common.CallBack)
	r.POST("/users", controllers.CreateUser)
	r.POST("/classrooms",controllers.CreateClassroom)
	r.POST("/classrooms/:classroomid/teacher/:teacheruuid", controllers.AssignClassTeacher)
	r.POST("/classrooms/:classroomid/student/:studentuuid", controllers.AssignClassStudent)
	r.GET("/users", controllers.GetUsers)
	r.GET("/classrooms", controllers.GetClassrooms)
	r.GET("/teachers", controllers.GetTeachers)
	r.GET("/students", controllers.GetStudents)
	r.GET("/users/:id", controllers.GetUserByID)
	r.GET("/classrooms/:id",controllers.GetClassroomByID)
	r.GET("/teachers/:teacheruuid", controllers.GetTeacherByID)
	r.GET("/students/:studentuuid", controllers.GetStudentByID)
	r.PUT("/users/roles/:id", controllers.UpdateUserRole)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.PUT("/users/mobile/:id", controllers.UpdateUserMobile)
	r.PUT("/classrooms/:id", controllers.UpdateClassroom)
	r.PUT("/teachers/:teacheruuid", controllers.UpdateTeacher)
	r.PUT("/students/:studentuuid", controllers.UpdateStudent)
	r.PUT("/classrooms/:id/oldteacher/:oldteacheruuid/newteacher/:newteacheruuid",controllers.UpdateClassroomTeacher)
	r.DELETE("/users/:id",controllers.DeleteUser)
	r.DELETE("/teachers/:teacheruuid", controllers.DeleteTeacher)
	r.DELETE("/students/:studentuuid", controllers.DeleteStudent)
	r.DELETE("/classrooms/:id", controllers.DeleteClassroom)
	r.DELETE("/classroom/students/:classroomid/:studentuuid", controllers.DeleteClassStudent)
	r.DELETE("/classrooms/teachers/:classroomid/:teacheruuid", controllers.DeleteClassTeacher)
	r.Run()
}

// {	uuid:1,
// 	first_name:"Manan",
// 	middle_name:"",
// 	last_name:"Mittal",
// 	mobile:"7017765638",
// 	email:"mananiitd2905@gmail.com",},

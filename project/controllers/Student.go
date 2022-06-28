package controllers

import (
	// "database/sql"
	// "errors"
	"net/http"
	"project/db"
	"project/models"
	"project/models/service"

	// "project/models/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateStudent(UserId uint, c *gin.Context) error {
	var newStudent models.Student
	newStudent.UserId = UserId
	if err := db.DbConn.Create(&newStudent).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newStudent)
		return err
	}
	return nil
}

func GetStudents(c *gin.Context) {
	// Get all records
	is_Authorized:=service.Authorize(c,models.GetUser)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var students []models.Student

	if err := db.DbConn.Find(&students).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, students)
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func GetStudentByID(c *gin.Context) {
	is_Authorized:=service.Authorize(c,models.GetUser)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	uuid, err := strconv.Atoi(c.Param("studentuuid"))
	if err != nil {
		return
	}

	student := models.Student{Uuid: uint(uuid)}
	if err := db.DbConn.Take(&student).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, student)
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	is_Authorized:=service.Authorize(c,models.UpdateUserData["student"])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	uuid, err := strconv.Atoi(c.Param("studentuuid"))
	if err != nil {
		return
	}
	student := models.Student{Uuid: uint(uuid)}
	if err:=db.DbConn.Take(&student).Error;err!=nil{
		c.IndentedJSON(http.StatusBadRequest,student)
		return
	}
	if err := UpdateUserById(c, student.UserId); err != nil {
		c.IndentedJSON(http.StatusBadRequest,student)
		return
	}
	c.IndentedJSON(http.StatusOK,student)
}

func DeleteStudent(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.DeleteUser["student"])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	studentId, err := strconv.Atoi(c.Param("studentuuid"))
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,studentId)
		return
	}
	student:=models.Student{Uuid: uint(studentId)}
	if err:=db.DbConn.Take(&student).Error;err!=nil{
		c.IndentedJSON(http.StatusNotFound,student)
		return
	}
	user:=models.User{Id: student.UserId}
	if err:=db.DbConn.Delete(&user).Error;err!=nil{
		c.IndentedJSON(http.StatusNotModified,user)
		return;
	}
	if err:=db.DbConn.Delete(&student).Error;err!=nil{
		c.IndentedJSON(http.StatusNotModified,student)
		return;
	}
	c.IndentedJSON(http.StatusOK,student)
}
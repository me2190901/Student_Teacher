package controllers

import (
	// "database/sql"
	"net/http"
	"project/db"
	"project/models"
	"project/models/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

func AssignClassStudent(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.AssignClassStudent)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var newClassStudent models.ClassroomStudents
	classroomId, err := strconv.Atoi(c.Param("classroomid"))
	if err!=nil{
		return
	}
	studentId, err := strconv.Atoi(c.Param("studentuuid"))
	if err!=nil{
		return
	}
	newClassStudent.ClassroomId = uint(classroomId)
	newClassStudent.StudentUuid = uint(studentId)
	if err := db.DbConn.Create(&newClassStudent).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassStudent)
		return
	}
	c.IndentedJSON(http.StatusCreated, newClassStudent)
}

func DeleteClassStudent(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.DeleteClassStudent)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var classStudent models.ClassroomStudents
	classroomId, err := strconv.Atoi(c.Param("classroomid"))
	if err!=nil{
		return
	}
	studentId, err := strconv.Atoi(c.Param("studentuuid"))
	if err!=nil{
		return
	}
	classStudent.ClassroomId = uint(classroomId)
	classStudent.StudentUuid = uint(studentId)
	if err := db.DbConn.Delete(&classStudent).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, classStudent)
		return
	}
	c.IndentedJSON(http.StatusCreated, classStudent)
}
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

func AssignClassTeacher(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.AssignClassTeacher)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var newClassTeacher models.ClassroomTeachers
	classroomId, err := strconv.Atoi(c.Param("classroomid"))
	if err!=nil{
		return
	}
	teacherId, err := strconv.Atoi(c.Param("teacheruuid"))
	if err!=nil{
		return
	}
	newClassTeacher.ClassroomId = uint(classroomId)
	newClassTeacher.TeacherUuid = uint(teacherId)
	if err := db.DbConn.Create(&newClassTeacher).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassTeacher)
		return
	}
	c.IndentedJSON(http.StatusCreated, newClassTeacher)
}

func UpdateClassroomTeacher(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.UpdateClassTeacher)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var newClassTeacher models.ClassroomTeachers
	classroomId, err := strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,newClassTeacher)
		return
	}
	oldteacherId, err := strconv.Atoi(c.Param("oldteacheruuid"))
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,newClassTeacher)
		return
	}
	newteacherId, err := strconv.Atoi(c.Param("newteacheruuid"))
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,newClassTeacher)
		return
	}
	classTeacher:=models.ClassroomTeachers{ClassroomId:uint(classroomId),TeacherUuid: uint(oldteacherId)}
	newClassTeacher.TeacherUuid=uint(newteacherId)

	if err := db.DbConn.Model(&classTeacher).Omit("ClassroomId").Updates(newClassTeacher).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassTeacher)
		return
	}
	c.IndentedJSON(http.StatusOK,newClassTeacher)
}

func DeleteClassTeacher(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.DeleteClassTeacher)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var classTeacher models.ClassroomTeachers
	classroomId, err := strconv.Atoi(c.Param("classroomid"))
	if err!=nil{
		return
	}
	teacherId, err := strconv.Atoi(c.Param("teacheruuid"))
	if err!=nil{
		return
	}
	classTeacher.ClassroomId = uint(classroomId)
	classTeacher.TeacherUuid = uint(teacherId)
	if err := db.DbConn.Delete(&classTeacher).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, classTeacher)
		return
	}
	c.IndentedJSON(http.StatusCreated, classTeacher)
}
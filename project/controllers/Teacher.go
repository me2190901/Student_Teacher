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

func CreateTeacher(UserId uint, c *gin.Context) error {
	var newTeacher models.Teacher
	newTeacher.UserId = UserId
	if err := db.DbConn.Create(&newTeacher).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newTeacher)
		return err
	}
	return nil
}

func GetTeachers(c *gin.Context) {
	// Get all records
	is_Authorized:=service.Authorize(c,models.GetUser)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var teachers []models.Teacher

	if err := db.DbConn.Find(&teachers).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, teachers)
		return
	}
	c.IndentedJSON(http.StatusOK, teachers)
}

func GetTeacherByID(c *gin.Context) {
	is_Authorized:=service.Authorize(c,models.GetUser)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	uuid, err := strconv.Atoi(c.Param("teacheruuid"))
	if err != nil {
		return
	}
	teacher := models.Teacher{Uuid: uint(uuid)}
	if err := db.DbConn.Take(&teacher).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, teacher)
		return
	}
	c.IndentedJSON(http.StatusOK, teacher)
}

func UpdateTeacher(c *gin.Context) {
	is_Authorized:=service.Authorize(c,models.UpdateUserData["teacher"])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	uuid, err := strconv.Atoi(c.Param("teacheruuid"))
	if err != nil {
		return
	}
	teacher := models.Teacher{Uuid: uint(uuid)}
	if err:=db.DbConn.Take(&teacher);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,teacher)
		return
	}
	if err := UpdateUserById(c, teacher.UserId); err != nil {
		c.IndentedJSON(http.StatusBadRequest,teacher)
		return
	}
	c.IndentedJSON(http.StatusOK,teacher)
}

func DeleteTeacher(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.DeleteUser["teacher"])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	teacherId, err := strconv.Atoi(c.Param("teacheruuid"))
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,teacherId)
		return
	}
	teacher:=models.Teacher{Uuid: uint(teacherId)}
	if err:=db.DbConn.Take(&teacher).Error;err!=nil{
		c.IndentedJSON(http.StatusNotFound,teacher)
		return
	}
	user:=models.User{Id: teacher.UserId}
	if err:=db.DbConn.Delete(&user).Error;err!=nil{
		c.IndentedJSON(http.StatusNotModified,user)
		return;
	}
	if err:=db.DbConn.Delete(&teacher).Error;err!=nil{
		c.IndentedJSON(http.StatusNotModified,teacher)
		return;
	}
	c.IndentedJSON(http.StatusOK,teacher)
}
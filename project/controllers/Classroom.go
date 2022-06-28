package controllers

import (
	// "database/sql"
	"errors"
	"net/http"
	"project/db"
	"project/models"
	"project/models/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func classroomConstraints(newClassroom models.Classroom) error{
	if(!StringNotNull(newClassroom.Name)){
		return errors.New("class name should not be null")
	}
	return nil
}

func CreateClassroom(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.CreateClassroom)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var newClassroom models.Classroom
	if err := c.BindJSON(&newClassroom); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassroom)
		return
	}
	if err:=classroomConstraints(newClassroom);err!=nil{
		c.IndentedJSON(http.StatusBadRequest, newClassroom)
		return
	}
	if err := db.DbConn.Create(&newClassroom).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassroom)
		return
	}
	c.IndentedJSON(http.StatusCreated, newClassroom)
}


func GetClassrooms(c *gin.Context) {
	// Get all records
	is_Authorized:=service.Authorize(c,models.GetClassroom)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	var classrooms []models.Classroom

	if err := db.DbConn.Find(&classrooms).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, classrooms)
		return
	}
	c.IndentedJSON(http.StatusOK, classrooms)
}

func UpdateClassroom(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.UpdateClassroom)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	uuid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	classroom := models.Classroom{Id: uint(uuid)}
	var newClassroom models.Classroom
	if err := c.ShouldBindBodyWith(&newClassroom, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassroom)
		return
	}
	if err := classroomConstraints(newClassroom); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassroom)
		return
	}
	if err := db.DbConn.Model(&classroom).Omit("Id").Updates(newClassroom).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newClassroom)
		return
	}
	c.IndentedJSON(http.StatusOK, classroom)
}

func GetClassroomByID(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.GetClassroom)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	uuid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	classroom := models.Classroom{Id: uint(uuid)}
	if err := db.DbConn.Take(&classroom).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, classroom)
		return
	}
	
	c.IndentedJSON(http.StatusOK, classroom)
}

func DeleteClassroom(c *gin.Context){
	is_Authorized:=service.Authorize(c,models.DeleteClassroom)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	classroomId, err := strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,classroomId)
		return
	}
	classroom:=models.Classroom{Id: uint(classroomId)}
	if err:=db.DbConn.Delete(&classroom).Error;err!=nil{
		c.IndentedJSON(http.StatusNotModified,classroom)
		return;
	}
	c.IndentedJSON(http.StatusOK,classroom)
}
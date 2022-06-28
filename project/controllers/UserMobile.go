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

func MobileConstraints(newUserMobile models.UserMobile) error{
	if(!StringNotNull(newUserMobile.Mobile)){
		return errors.New("mobile number should not be null")
	}else if(len(newUserMobile.Mobile)!=10){
		return errors.New("invalid Mobile number")
	}
	return nil
}

func CreateUserMobile(newUserMobile models.UserMobile, c *gin.Context) error {
	if err := db.DbConn.Create(&newUserMobile).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUserMobile)
		return err
	}
	c.IndentedJSON(http.StatusCreated,newUserMobile)
	return nil
}

func UpdateUserMobile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}
	userMobile := models.UserMobile{UserId: uint(id)}
	var newUserMobile models.UserMobile
	if err := c.ShouldBindBodyWith(&newUserMobile, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUserMobile)
		return
	}
	user:=models.User{Id: newUserMobile.UserId}
	if err:=db.DbConn.Take(&user).Error;err!=nil{
		c.IndentedJSON(http.StatusBadRequest, "User not found")
	}
	is_Authorized:=service.Authorize(c,models.UpdateUserMobile[user.Role])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}

	if err := db.DbConn.Model(&userMobile).Select("Mobile").Updates(newUserMobile).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUserMobile)
		return
	}
	

	c.IndentedJSON(http.StatusOK, userMobile)
}
package controllers

import (
	// "database/sql"
	"errors"
	"net/http"
	"net/mail"
	"project/db"
	"project/models"
	"project/models/service"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func StringNotNull(s string) bool {
	return len(s) != 0
}

func user_Constraints(newUser models.User) error {
	if !StringNotNull(newUser.First_Name) || !StringNotNull(newUser.Last_Name) || !StringNotNull(newUser.Email) {
		return errors.New("cannot create user with specified details")
	} else if _, err := mail.ParseAddress(newUser.Email); err != nil {
		return errors.New("cannot create user with specified email")
	} else if !(models.Roles[newUser.Role]) {
		return errors.New("invalid Role Type")
	}
	return nil
}


func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}
	is_Authorized:=service.Authorize(c,models.CreateUser[newUser.Role])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	if err := user_Constraints(newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}
	var newUserMobile models.UserMobile
	newUserMobile.UserId = newUser.Id
	if err := c.ShouldBindBodyWith(&newUserMobile, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}
	if err := MobileConstraints(newUserMobile); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}

	if err := db.DbConn.Create(&newUser).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}
	if err := CreateUserMobile(newUserMobile, c); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}

	if newUser.Role == "teacher" {
		if err := CreateTeacher(newUser.Id, c); err != nil {
			c.IndentedJSON(http.StatusBadRequest, newUser)
			return
		}
	} else if newUser.Role == "student" {
		if err := CreateStudent(newUser.Id, c); err != nil {
			c.IndentedJSON(http.StatusBadRequest, newUser)
			return
		}
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUsers(c *gin.Context) {
	// Get all records
	
	var users []models.User
	is_Authorized:=service.Authorize(c,models.GetUser)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	if err := db.DbConn.Find(&users).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, users)
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	// Get all records
	is_Authorized:=service.Authorize(c,models.GetUser)
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	uuid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, uuid)
		return
	}

	user := models.User{Id: uint(uuid)}
	if err := db.DbConn.Take(&user).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, user)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func UpdateUserRole(c *gin.Context) {
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	user := models.User{Id: uint(id)}
	var newUser models.User
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}
	if err:= db.DbConn.Take(&user).Error;err!=nil{
		c.IndentedJSON(http.StatusBadRequest,"user not found")
		return
	}

	is_Authorized1:=service.Authorize(c,models.UpdateUserRole[user.Role])
	if(!is_Authorized1){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	if user.Role=="student"{

		student:=models.Student{UserId: user.Id }
		db.DbConn.Take(&student)
		db.DbConn.Delete(&student)
	}else if user.Role=="teacher"{
		teacher:=models.Teacher{UserId: user.Id }
		db.DbConn.Take(&teacher)
		db.DbConn.Delete(&teacher)
	}

	is_Authorized2:=service.Authorize(c,models.UpdateUserRole[newUser.Role])
	if(!is_Authorized2){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	if newUser.Role=="student"{
		if err:=CreateStudent(user.Id,c);err!=nil{
			c.IndentedJSON(http.StatusBadRequest, "Cannot create student with this details")
			return
		}
	}else if newUser.Role=="teacher"{
		if err:=CreateTeacher(user.Id,c);err!=nil{
			c.IndentedJSON(http.StatusBadRequest, "Cannot create student with this details")
			return
		}
	}
	if models.Roles[newUser.Role] {
		if err := db.DbConn.Model(&user).Select("Role").Updates(newUser).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, newUser)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	user := models.User{Id: uint(id)}
	if err := db.DbConn.Find(&user).Error; err != nil {
		return
	}
	is_Authorized:=service.Authorize(c,models.UpdateUserData[user.Role])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}
	if err := UpdateUserById(c, uint(id)); err != nil {
		return
	}
}

func UpdateUserById(c *gin.Context, id uint) error {
	user := models.User{Id: uint(id)}
	if err := db.DbConn.Find(&user).Error; err != nil {
		return err
	}
	var newUser models.User
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return err
	}
	if err := user_Constraints(newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return err
	}
	if err := db.DbConn.Model(&user).Omit("Role", "Id").Updates(newUser).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return err
	}
	c.IndentedJSON(http.StatusOK, user)
	return nil
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, userId)
		return
	}
	user := models.User{Id: uint(userId)}
	if err:=db.DbConn.Take(&user).Error;err!=nil{
		c.IndentedJSON(http.StatusBadRequest, "User does not exist")
		return
	}
	is_Authorized:=service.Authorize(c,models.DeleteUser[user.Role])
	if(!is_Authorized){
		c.IndentedJSON(http.StatusBadRequest, "You are not authorized to do this call")
		return
	}

	if err := db.DbConn.Delete(&user).Error; err != nil {
		c.IndentedJSON(http.StatusNotModified, user)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

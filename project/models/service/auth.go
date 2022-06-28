package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)
func contains(list []string,s string) bool{
	for i:=0;i<len(list);i++{
		if list[i]==s{
			return true
		}
	}
	return false
}
func Validate(c *gin.Context) (string,error){
	access_token:=c.GetHeader("Authorization")
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + access_token)
	if err != nil {
		fmt.Println("User Data fetch failed")
		return "",err
	}
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Json Parsing Failed")
		return "",err
	}
	var data models.UserTokenData
	json.Unmarshal(userData,&data)
	user:=models.User{
		Email: data.Email,
	}
	if err:=db.DbConn.Where("Email=? ",user.Email).First(&user).Error;err!=nil{
		return "",err
	}
	if user.Token!=access_token{
		return "",errors.New("token mismatch")
	}
	return user.Role,nil
}

func Authorize(c *gin.Context,access_list []string) bool{
	role,err:=Validate(c)
	if err!=nil || !contains(access_list,role){
		return false
	}
	return true
}
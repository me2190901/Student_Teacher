package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"project/db"
	"project/models"
)

func UpdateToken(token string) error{
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return err
	}
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var data models.UserTokenData
	json.Unmarshal(userData,&data)
	if(!data.Verify){
		return errors.New("user not verified ")
	}
	user:=models.User{
		Email: data.Email,
	}
	if err:=db.DbConn.Where("Email=? ",user.Email).First(&user).Error;err!=nil{
		return err
	}
	user.Token=token
	db.DbConn.Save(&user)
	return nil
}
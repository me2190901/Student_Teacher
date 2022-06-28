package models

type User struct {
	Id          uint   `json:"uuid" gorm:"primary_key"`
	First_Name  string `json:"first_name" gorm:"type:VARCHAR(30);not null"`
	Middle_Name string `json:"middle_name" gorm:"size:30;"`
	Last_Name   string `json:"last_name" gorm:"type:VARCHAR(30);not null"`
	Role        string `json:"role" sql:"type:role_access" gorm:"not null"`
	Email       string `json:"email" gorm:"not null;size:100;index;unique"`
	Token		string `json:"token"`
}
package models

type UserMobile struct {
	Mobile string `json:"mobile" gorm:"size:10;primary_key"`
	UserId uint   `json:"user_Id" gorm:"index;primary_key"`
	User   User   `gorm:"foreign_key:UderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
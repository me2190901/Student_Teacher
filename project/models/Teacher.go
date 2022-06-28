package models

type Teacher struct {
	Uuid   uint `json:"uuid" gorm:"primary_key;index"`
	UserId uint `json:"user_id" gorm:"unique,not null"`
	User   User `gorm:"foreign_key:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

}

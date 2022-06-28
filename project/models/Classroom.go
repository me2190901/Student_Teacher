package models

type Classroom struct {
	Id          uint    `json:"id" gorm:"primary_key"`
	Name		string 	`json:"name" gorm:"unique;size:30;not null"`
}

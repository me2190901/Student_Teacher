package models

type ClassroomStudents struct {
	ClassroomId uint `json:"id" gorm:"primary_key"`
	StudentUuid uint `json:"Student_uuid" gorm:"primary_key"`
	Student     Student `gorm:"foreign_key:StudentUuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Classroom	Classroom	`gorm:"foreign_key:ClassroomId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

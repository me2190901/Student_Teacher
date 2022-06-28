package models

type ClassroomTeachers struct {
	ClassroomId uint `json:"id" gorm:"primary_key"`
	TeacherUuid uint `json:"teacher_uuid" gorm:"primary_key"`
	Teacher     Teacher `gorm:"foreign_key:TeacherUuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Classroom	Classroom	`gorm:"foreign_key:ClassroomId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

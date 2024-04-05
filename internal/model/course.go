package model

import (
	"time"
)

type Course struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	CourseID  string    `gorm:"uniqueIndex;type:varchar(100)" json:"courseId"`
	Name      string    `json:"name"`
	Lessons   int32     `json:"lessons"`
}

func (m *Course) TableName() string {
	return "courses"
}

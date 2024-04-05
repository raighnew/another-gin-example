package model

import "time"

type Enrollment struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	StudentID string    `gorm:"index:idx_enrollment" json:"studentId"`
	CourseID  string    `gorm:"index:idx_enrollment" json:"courseId"`
}

func (m *Enrollment) TableName() string {
	return "enrollments"
}

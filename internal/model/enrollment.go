package model

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
}

func (m *Enrollment) TableName() string {
    return "enrollment"
}

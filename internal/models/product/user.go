package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AuthID uint `gorm:"unique;not null"`
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);not null"`
	Email string `gorm:"varchar(40);not null;unique;"`
	Password string `gorm:"size 255;not null"`
}

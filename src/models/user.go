package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
	UpdatedAt string `gorm:"not null"`
	CreatedBy string `gorm:"not null"`
}

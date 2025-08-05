package models

import ("gorm.io/gorm")	

type User struct {
	gorm.Model //id, created_at, updated_at, deleted_at
	Name string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}
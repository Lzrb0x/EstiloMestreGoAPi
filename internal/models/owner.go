package models

import (
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model // id, created_at, updated_at, deleted_at

	UserID uint `gorm:"not null;unique"`
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
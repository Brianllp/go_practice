package models

import (
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

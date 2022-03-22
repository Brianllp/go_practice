package models

import (
	"github.com/Brianllp/go_practice/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUsers() (users []User) {
	db := database.GetDB()
	db.Find(&users)
	return users
}

func FindUserByID(id string) (user User) {
	db := database.GetDB()
	db.Where("id = ?", id).First(&user)
	return user
}

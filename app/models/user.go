package models

import (
	"github.com/Brianllp/go_practice/app/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUsers() (users []User) {
	database.DB.Find(&users)
	return users
}

func FindUserByID(id string) (user User) {
	database.DB.Where("id = ?", id).First(&user)
	return user
}

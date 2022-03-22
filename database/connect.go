package database

import (
	"fmt"
	"os"

	godotnev "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	// ENV読み取り
	err_env := godotnev.Load(".env")
	if err_env != nil {
		fmt.Printf("読み込みに失敗しました: %v", err_env)
	}

	// 接続情報を設定
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := "tcp(db:3306)"
	DBNAME := os.Getenv("DB_NAME")

	dsn := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"

	// DBに接続
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	}

	return db, err
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

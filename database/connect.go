package database

import (
	"fmt"
	"os"

	// "github.com/Brianllp/go_practice/models"
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

	// InitDB()

	return db, err
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// func InitDB() {
// 	DB.AutoMigrate(&models.User{})
// }

// Todo: migrateとかは後で考える
// func InitDB() {
// 	// ENV読み取り
// 	err := godotnev.Load(".env")
// 	if err != nil {
// 		fmt.Printf("読み込みに失敗しました: %v", err)
// 	}

// 	// 接続情報を設定
// 	// USER := os.Getenv("DB_USER")
// 	// PASS := os.Getenv("DB_PASS")
// 	USER := "webuser"
// 	PASS := "webpass"
// 	HOST := "tcp(db:3306)"
// 	DBNAME := "go_development"

// 	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"

// 	// DBに接続
// 	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	sqlDB, db_err := db.DB()
// 	if db_err != nil {
// 		panic(db_err.Error())
// 	}
// 	defer sqlDB.Close()

// 	// Userテーブルの作成、マイグレーション
// 	db.AutoMigrate(&User{})

// 	// Userデータ取得
// 	var users []User
// 	result := db.Find(&users)

// 	// Userデータがなければ作成する
// 	user_data_count := result.RowsAffected
// 	if user_data_count == 0 {
// 		user := User{Name: "HogeHoge_test", Age: 40}
// 		created_result := db.Create(&user)

// 		if result.Error != nil {
// 			panic(err.Error())
// 		}

// 		println(created_result.RowsAffected)
// 		fmt.Println(created_result)
// 		return
// 	}

// 	fmt.Println(users)
// }

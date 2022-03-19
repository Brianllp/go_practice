package main

import (
	"fmt"
	"net/http"
	"os"

	godotnev "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// ENV読み取り
	err := godotnev.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}

	// 接続情報を設定
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := "tcp(db:3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"

	// DBに接続
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, db_err := db.DB()
	if db_err != nil {
		panic(db_err.Error())
	}
	defer sqlDB.Close()

	// Userテーブルの作成、マイグレーション
	db.AutoMigrate(&User{})

	// Userデータ取得
	var users []User
	result := db.Find(&users)

	// Userデータがなければ作成する
	user_data_count := result.RowsAffected
	if user_data_count == 0 {
		user := User{Name: "HogeHoge_test", Age: 40}
		created_result := db.Create(&user)

		if result.Error != nil {
			panic(err.Error())
		}

		println(created_result.RowsAffected)
		fmt.Println(created_result)
		return
	}

	fmt.Println(users)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":3030"))
}

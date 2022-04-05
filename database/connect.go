package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB(is_root bool) (*gorm.DB, error) {
	var (
		USER   string
		PASS   string
		HOST   = "tcp(" + os.Getenv("CONTAINER_NAME") + ":" + os.Getenv("DB_PORT") + ")"
		DBNAME = os.Getenv("DB_NAME")
	)

	// 接続情報を設定
	if is_root {
		USER = os.Getenv("MYSQL_ROOT_USER")
		PASS = os.Getenv("MYSQL_ROOT_PASSWORD")
	} else {
		USER = os.Getenv("DB_USER")
		PASS = os.Getenv("DB_PASS")
	}

	dsn := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"

	fmt.Printf("dsn is: %s \n", dsn)
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

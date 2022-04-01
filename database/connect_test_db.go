package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var test_db *gorm.DB

var (
	USER   = os.Getenv("DB_USER")
	PASS   = os.Getenv("DB_PASS")
	HOST   = "tcp(" + os.Getenv("CONTAINER_NAME") + ":" + os.Getenv("DB_PORT") + ")"
	DBNAME = os.Getenv("TEST_DB_NAME")
)

func CreateTestDB() {
	db, _ := ConnectDB(true)
	defer CloseDB()

	db.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME)
	// テスト用DBの操作をroot以外のユーザーで行うために権限付与する
	db.Exec("GRANT ALL ON " + DBNAME + ".* TO " + USER)
}

func ConnectTestDB() (*gorm.DB, error) {
	dsn := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"
	fmt.Printf("test db dsn is: %s \n", dsn)

	// DBに接続
	var err error
	test_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	return test_db, err
}

func GetTestDB() *gorm.DB {
	return test_db
}

func CloseTestDB() {
	sqlDB, _ := test_db.DB()
	sqlDB.Close()
}

package test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
)

var db *sql.DB
var (
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbName   = os.Getenv("DB_NAME")
	port     = "3306"
	host     = "tcp(db:3306)"
	dsn      = user + ":" + password + "@" + host + "/" + dbName + "?parseTime=true"
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	// pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	fmt.Println("---------------1----------")

	// buildOptions := &dockertest.BuildOptions{
	// 	// for m1 mac
	// 	Platform: "linux/x86_64",
	// }

	// // 	// Dockerコンテナ起動時の細かいオプションを指定する
	// runOptions := &dockertest.RunOptions{
	// 	Repository: "mysql",
	// 	Tag:        "8.0",
	// 	Env: []string{
	// 		"MYSQL_ROOT_PASSWORD=secret",
	// 		// "MYSQL_USER=" + user,
	// 		// "MYSQL_PASSWORD=" + password,
	// 		// "MYSQL_DATABASE=" + dbName,
	// 	},
	// }

	// pulls an image, creates a container based on it and runs it
	// resource, err := pool.BuildAndRunWithBuildOptions(buildOptions, runOptions)
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	fmt.Println("=========================")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

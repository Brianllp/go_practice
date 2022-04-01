package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/Brianllp/go_practice/database"
)

func TestMain(m *testing.M) {
	database.CreateTestDB()
	database.ConnectTestDB()
	defer database.CloseDB()

	fmt.Println("before run")

	code := m.Run()

	fmt.Println("after run")

	os.Exit(code)
}

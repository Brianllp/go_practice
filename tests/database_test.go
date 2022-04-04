package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/Brianllp/go_practice/database"
	"github.com/Brianllp/go_practice/models"
)

func TestMain(m *testing.M) {
	database.CreateTestDB()
	database.ConnectTestDB()
	models.Migration(database.GetTestDB())
	defer database.CloseDB()

	fmt.Println("before run")

	code := m.Run()

	fmt.Println("after run")

	database.DropTestDB()

	os.Exit(code)
}

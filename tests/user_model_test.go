package tests

import (
	"fmt"
	"testing"

	"github.com/Brianllp/go_practice/database"
	"github.com/Brianllp/go_practice/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var user models.User

func TestFindUserByID(t *testing.T) {
	user1 := models.User{Name: "HogeHoge_test", Age: 45}

	db, err := database.ConnectTestDB()
	if err != nil {
		fmt.Println(err)
	}
	defer database.CloseTestDB()
	// トランザクションを貼る
	tx := db.Begin()

	if err := tx.Error; err != nil {
		fmt.Println(err)
	}

	// ユーザーデータの作成
	CreateUser(tx)

	created_user := models.FindUserByID(tx, "1")

	// Todo: ルーティングのテスト
	// e := router.NewRouter()
	// req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	// rec := httptest.NewRecorder()

	// e.ServeHTTP(rec, req)

	// created_user := rec.Body.String()
	// fmt.Println(created_user)

	// assert.Equal(t, http.StatusOK, rec.Code)

	assert.Equal(t, user1.Name, created_user.Name)

	// rollbackで生成したデータを削除
	tx.Rollback()
}

// Todo: model に user の create 処理を実装する
func CreateUser(tx *gorm.DB) {
	user = models.User{Name: "HogeHoge_test", Age: 40}
	tx.Create(&user)
}

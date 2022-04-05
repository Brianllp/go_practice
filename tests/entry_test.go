package tests

import (
	"fmt"
	"testing"

	"github.com/Brianllp/go_practice/database"
	"github.com/Brianllp/go_practice/models"
	"github.com/stretchr/testify/assert"
)

var entry_data = models.Entry{
	UUID:  "test_uuid",
	Title: "test_title",
	Body:  "test_Body",
}

var entry_data2 = models.Entry{
	UUID:  "test_uuid",
	Title: "test_title2",
	Body:  "test_Body2",
}

var entry_data3 = []models.Entry{
	{
		UUID:  "test_uuid3",
		Title: "test_title3",
		Body:  "test_Body3",
	},
	{
		UUID:  "test_uuid4",
		Title: "test_title4",
		Body:  "test_Body4",
	},
}

func TestCreateEntry(t *testing.T) {
	var created_entry models.Entry

	db, err := database.ConnectTestDB()
	if err != nil {
		fmt.Println(err)
	}
	defer database.CloseTestDB()

	tx := db.Begin()
	if err := tx.Error; err != nil {
		fmt.Println(err)
	}

	models.CreateOrUpdateEntry(tx, entry_data)

	tx.Debug().Where("uuid = ?", entry_data.UUID).First(&created_entry)

	assert.Equal(t, entry_data.UUID, created_entry.UUID)
	assert.Equal(t, entry_data.Title, created_entry.Title)
	assert.Equal(t, entry_data.Body, created_entry.Body)

	tx.Rollback()
}

func TestUpdateEntry(t *testing.T) {
	var created_entry models.Entry

	db, err := database.ConnectTestDB()
	if err != nil {
		fmt.Println(err)
	}
	defer database.CloseTestDB()

	tx := db.Begin()
	if err := tx.Error; err != nil {
		fmt.Println(err)
	}

	models.CreateOrUpdateEntry(tx, entry_data)  // create entry
	models.CreateOrUpdateEntry(tx, entry_data2) // update entry

	tx.Debug().Where("uuid = ?", entry_data.UUID).First(&created_entry)

	assert.Equal(t, entry_data2.Title, created_entry.Title)
	assert.Equal(t, entry_data2.Body, created_entry.Body)

	tx.Rollback()
}

func TestIndexEntries(t *testing.T) {
	var test_data_UUIDs []string
	var test_data_Titles []string
	var test_data_Bodies []string

	var created_data_UUIDs []string
	var created_data_Titles []string
	var created_data_Bodies []string

	for _, test_data := range entry_data3 {
		test_data_UUIDs = append(test_data_UUIDs, test_data.UUID)
		test_data_Titles = append(test_data_Titles, test_data.Title)
		test_data_Bodies = append(test_data_Bodies, test_data.Body)
	}

	db, err := database.ConnectTestDB()
	if err != nil {
		fmt.Println(err)
	}
	defer database.CloseTestDB()

	tx := db.Begin()
	if err := tx.Error; err != nil {
		fmt.Println(err)
	}

	// create multi entries
	for _, entry := range entry_data3 {
		models.CreateOrUpdateEntry(tx, entry)
	}

	entries := models.IndexEntries(tx)

	for _, created_data := range entries {
		created_data_UUIDs = append(created_data_UUIDs, created_data.UUID)
		created_data_Titles = append(created_data_Titles, created_data.Title)
		created_data_Bodies = append(created_data_Bodies, created_data.Body)
	}

	assert.Equal(t, test_data_UUIDs, created_data_UUIDs)
	assert.Equal(t, test_data_Titles, created_data_Titles)
	assert.Equal(t, test_data_Bodies, created_data_Bodies)

	tx.Rollback()
}

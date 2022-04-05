package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	ID        uint      `gorm:"primaryKey;unique;not null;autoIncrement" json:"id"`
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func IndexEntries(db *gorm.DB) (entry []Entry) {
	db.Find(&entry)
	return entry
}

func CreateOrUpdateEntry(db *gorm.DB, entry Entry) {
	resultData := Entry{}
	record := db.First(&resultData, "uuid = ?", entry.UUID)
	isNotExistEntryRecord := errors.Is(record.Error, gorm.ErrRecordNotFound)

	if isNotExistEntryRecord {
		db.Create(&entry)
	} else {
		resultData.Title = entry.Title
		resultData.Body = entry.Body

		db.Save(&resultData)
	}
}

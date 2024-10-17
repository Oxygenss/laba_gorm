package repository

import (
	"github.com/pkg/errors"

	"github.com/Oxygenss/laba_gorm/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "init db")
	}

	if err := db.AutoMigrate(&models.Students{}); err != nil {
		return nil, errors.Wrap(err, "failed to auto migrate")
	}
	return db, nil
}

func Connection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "init db")
	}

	return db, nil
}

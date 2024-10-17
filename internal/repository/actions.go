package repository

import (
	"github.com/Oxygenss/laba_gorm/internal/models"
	"github.com/pkg/errors"
)

func GetAll() ([]models.Students, error) {
	db, err := Connection()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connection database")
	}

	var students []models.Students
	db.Find(&students)
	return students, nil
}

func CreateStudent(student *models.Students) error {
	db, err := Connection()
	if err != nil {
		return errors.Wrap(err, "failed to connection database")
	}

	err = db.Create(&student).Error
	if err != nil {
		return errors.Wrap(err, "failed to create student")
	}

	return nil
}

func DeleteStudent(id int) error {
	db, err := Connection()
	if err != nil {
		return errors.Wrap(err, "failed to connection database")
	}

	err = db.Where("id = ?", id).Delete(&models.Students{}).Error
	if err != nil {
		return errors.Wrap(err, "failed to delete student")
	}

	return nil
}

package repository

import (
	"time"

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

// func FilterStudents(minAge, maxAge int, before, after string) ([]models.Students, error) {
// 	db, err := Connection()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to connect to database")
// 	}

// 	var students []models.Students
// 	query := db.Model(&models.Students{})

// 	if minAge > 0 {
// 		//query = query.Where("age >= ?", minAge)
// 	}
// 	if maxAge > 0 {
// 		//query = query.Where("age <= ?", maxAge)
// 	}

// 	if before != "" {
// 		beforeDate, err := time.Parse("2006-01-02", before)
// 		if err != nil {
// 			return nil, errors.Wrap(err, "invalid 'before' date format")
// 		}
// 		query = query.Where("start_day <= ?", beforeDate)
// 	}

// 	if after != "" {
// 		afterDate, err := time.Parse("2006-01-02", after)
// 		if err != nil {
// 			return nil, errors.Wrap(err, "invalid 'after' date format")
// 		}
// 		query = query.Where("start_day >= ?", afterDate)
// 	}

// 	err = query.Find(&students).Error
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to filter students")
// 	}

// 	return students, nil
// }

func CalculateAge(birthday time.Time) int {
	now := time.Now()
	age := now.Year() - birthday.Year()

	if now.YearDay() < birthday.YearDay() {
		age--
	}

	return age
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

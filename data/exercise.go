package data

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Exercise struct {
	gorm.Model
	UID string `gorm:"not null"`
	User User `gorm:"foreignKey:UID;association_foreignkey:uid'"`
	Description string
	Duration int
	Date time.Time
}

func AddExercise(exercise *Exercise) (err error){
	err = Db.Create(&exercise).Error
	if err != nil {
		log.Print(err)
		return err
	}
	return
}

func GetUserExercises(uid string) ([]Exercise, error) {
	exercises := []Exercise{}
	rows, err := Db.Model(&Exercise{}).Where("uid = ?", uid).Select("uid, description, duration, date").Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return exercises, err
	}
	for rows.Next() {
		var ex Exercise
		Db.ScanRows(rows, &ex)
		exercises = append(exercises, ex)
	}
	return exercises, err
}

func GetUserExercisesQuery(uid string, fromdate time.Time, todate time.Time, limit int) ([]Exercise, error) {
	exercises := []Exercise{}
	if todate.IsZero() {
		todate = time.Now()
	}
	if limit < 1 {
		//Arbitrary number that should be queried.
		limit = 100
	}
	rows, err := Db.Model(&Exercise{}).Limit(limit).Where("uid = ? AND date BETWEEN ? AND ?", uid, fromdate, todate).Select("uid, description, duration, date").Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return exercises, err
	}
	for rows.Next() {
		var ex Exercise
		Db.ScanRows(rows, &ex)
		exercises = append(exercises, ex)
	}
	return exercises, err
}
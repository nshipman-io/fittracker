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

func GetUserExercises(uid string) ([]Exercise) {
	exercises := []Exercise{}
	//rows, err := Db.Raw("select uid,description, duration, date FROM exercises where uid = ?", uid).Rows()
	rows, err := Db.Model(&Exercise{}).Where("uid = ?", uid).Select("uid, description, duration, date").Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var ex Exercise
		Db.ScanRows(rows, &ex)
		exercises = append(exercises, ex)
	}
	return exercises
}
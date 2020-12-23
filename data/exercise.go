package data

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Exercise struct {
	gorm.Model
	Id int
	UserId string
	Description string
	Duration int
	Date string
}

func AddExercise(uid uuid.UUID, exercise *Exercise) (err error){
	user := User{}
	err = Db.First(&user, uid).Error
	if err != nil {
		return err
	}
	err = Db.Create(&exercise).Error
	user.Log = append(user.Log, *exercise)
	return
}
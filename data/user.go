package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	helper "github.com/nshipman-io/fittracker/helper"
	"github.com/satori/go.uuid"
	"log"
)

var Db *gorm.DB
func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=fittracker dbname=fittracker password=notsosafe sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&User{}, &Exercise{})
	Db.Model(&Exercise{}).AddForeignKey("uid","users(uid)", "ALLOW", "ALLOW")
}

type User struct {
	gorm.Model
	UID	string	`gorm:"unique;not null;primary key"`
	Username string	`gorm:"unique;not null"`
	Excercises	[]Exercise
}

func AddUser(user *User) (err error) {
	ruid := uuid.NewV4()
	uid:= helper.RemoveHyphenUIDString(ruid)
	user.UID = uid
	err = Db.Create(&user).Error
	if err != nil {
		return err
	}
	return
}

func GetUser(uid string) (*User, error) {
	user := User{}
	err := Db.Raw("SELECT * FROM users WHERE uid = ?",uid).Scan(&user).Error
	if err != nil {
		log.Println(err)
	}
	return &user, err
}

func GetAllUsers() ([]User, error) {
	users := []User{}
	result := Db.Find(&users)
	err := result.Error
	return users, err
}

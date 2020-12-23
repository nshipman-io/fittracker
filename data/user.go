package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"github.com/satori/go.uuid"
)

var Db *gorm.DB
func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=fittracker dbname=fittracker password=notsosafe sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&User{}, &Exercise{})
}

type User struct {
	gorm.Model
	UID	uuid.UUID
	Username string
	Log	[]Exercise
}

func genUUIDv4() (uid uuid.UUID, err error) {
	uid, err = uuid.NewV4()
	return uid, err
}

func AddUser(user *User) (err error) {
	uid, err := genUUIDv4()
	user.UID = uid
	if err != nil {
		return err
	}
	err = Db.Create(&user).Error
	if err != nil {
		return err
	}
	return
}

func GetUser(uid uuid.UUID) (*User, error) {
	user := User{}
	err := Db.First(&user, uid).Error
	return &user, err
}

func GetAllUsers() ([]User, error) {
	users := []User{}
	result := Db.Find(&users)
	err := result.Error
	return users, err
}
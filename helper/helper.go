package helper

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
	"time"
)

var timeLayout = "2006-01-02"

func ConvertTime(dt string) (time.Time, error) {
	t, err := time.Parse(timeLayout, dt)
	if err != nil {
		log.Println(err)
	}
	return t, err
}

func RemoveHyphenUIDString(uid uuid.UUID) (string) {
	suid := strings.Replace(uid.String(), "-","",-1)
	return suid
}
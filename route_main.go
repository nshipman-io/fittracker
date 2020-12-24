package main

import (
	"encoding/json"
	"fmt"
	"github.com/nshipman-io/fittracker/data"
	helper "github.com/nshipman-io/fittracker/helper"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserResp struct {
	Username string
	UID	string
}

type ExerciseResp struct {
	UID string
	Description string
	Duration int
	Date time.Time
}
func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "index","")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	user := data.User{}
	user.Username = r.Form.Get("username")
	err = data.AddUser(&user)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Username: %s already exists", user.Username)
		return
	}

	json, err := json.MarshalIndent(&user,"","\t\t")
	if err != nil {
		log.Println("Error Marshalling json in NewUser function")
	}
	w.Write(json)
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users,err := data.GetAllUsers()
	var uresp []UserResp
	if err != nil {
		fmt.Fprintf(w, "Could not retrieve users: %t", err)
	}
	for _, user := range(users) {
		ur := UserResp{
			Username: user.Username,
			UID: user.UID,
		}
		uresp = append(uresp, ur)
	}

	json, err := json.Marshal(&uresp)
	if err != nil {
		log.Println("Error Marshalling json in GetUsers function")
	}
	w.Write(json)
	return
}

func AddExercise(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "Error parsing form: ", err)
		return
	}

	dresp := r.Form.Get("duration")
	duration, err := strconv.Atoi(dresp)
	if err != nil {
		log.Println(err)
	}

	dtresp := r.Form.Get("date")
	date, err := helper.ConvertTime(dtresp)

	uid := r.Form.Get("userId")
	user,err := data.GetUser(uid)
	if err != nil {
		log.Println(err)
	}
	exercise := data.Exercise{
		UID: uid,
		User: *user,
		Description: r.Form.Get("description"),
		Duration: duration,
		Date: date,
	}

	err = data.AddExercise(&exercise)

	if err != nil {
		log.Println(err)
		return
	}


}
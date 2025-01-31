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
)

type UserResp struct {
	Username string
	UID	string
}

type LogResp struct {
	UID string
	Exercises []*ExerciseResp
}
type ExerciseResp struct {
	Description string
	Duration int
	Date string
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
	date, err := helper.ConvertStringToTime(dtresp)

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

func GetExerciseLog(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	fromdate := r.FormValue("from")
	todate := r.FormValue("to")
	limstr := r.FormValue("limit")
	ulog := LogResp{
		UID: uid,
	}

	if fromdate == "" && todate == "" && limstr == "" {
		exercises, err := data.GetUserExercises(uid)
		if err != nil {
			log.Println(err)
			return
		}
		for _,exercise := range exercises {
			date := helper.ConvertTimeToString(exercise.Date)
			exresp := ExerciseResp{
				Description: exercise.Description,
				Duration: exercise.Duration,
				Date: date,
			}
			ulog.Exercises = append(ulog.Exercises, &exresp)
		}

	}else {
		fdate, err := helper.ConvertStringToTime(fromdate)
		if err != nil {
			log.Println(err)
		}
		tdate, err := helper.ConvertStringToTime(todate)
		if err != nil {
			log.Println(err)
		}
		limit, err := strconv.Atoi(limstr)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, err)
			return
		}
		exercises, err := data.GetUserExercisesQuery(uid,fdate,tdate,limit)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, err)
			return
		}

		for _,exercise := range exercises {
			date := helper.ConvertTimeToString(exercise.Date)
			exresp := ExerciseResp{
				Description: exercise.Description,
				Duration: exercise.Duration,
				Date: date,
			}
			ulog.Exercises = append(ulog.Exercises, &exresp)
		}

	}
	json, err := json.MarshalIndent(&ulog,"","\t\t")
	if err != nil {
		log.Println("Error Marshalling json in GetExerciseLog function")
	}
	w.Write(json)
	return
}
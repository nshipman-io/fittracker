package main

import (
	"encoding/json"
	"fmt"
	"github.com/nshipman-io/fittracker/data"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"log"
)

type UserResp struct {
	Username string
	UID	uuid.UUID
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

	json, err := json.MarshalIndent(&uresp,"","\t\t")
	if err != nil {
		log.Println("Error Marshalling json in GetUsers function")
	}
	w.Write(json)
	return
}
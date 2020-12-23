package main

import (
	"encoding/json"
	"fmt"
	"github.com/nshipman-io/fittracker/data"
	"html/template"
	"net/http"
	"log"
)

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
		return
	}
	w.Write(json)
	return
}
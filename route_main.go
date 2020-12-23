package main

import (
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
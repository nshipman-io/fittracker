package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",fs))
	router.HandleFunc("/", Home)
	router.HandleFunc("/api/exercise/new-user", NewUser)
	router.HandleFunc("/api/exercise/users", GetUsers)
	router.HandleFunc("/api/exercise/add", AddExercise)
	router.Path("/api/exercise/log").
		HandlerFunc(GetExerciseLog).
		Queries("uid","{uid}","from","{from}", "to", "{to}","limit","{limit}").
		Name("GetExerciseLog")
	router.Path("/api/exercise/log").HandlerFunc(GetExerciseLog)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler:  router,
	}

	log.Println("Starting server at addr:", server.Addr)
	log.Fatal(server.ListenAndServe())
}
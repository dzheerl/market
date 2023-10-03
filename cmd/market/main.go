package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = ":3000"
)

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/send", submit)
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Println("Start server ")

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

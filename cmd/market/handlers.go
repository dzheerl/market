package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type LoginData struct {
	Form_email    string `json:"form_email"`
	Form_password string `json:"form_password"`
}

func login(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("pages/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func submit(w http.ResponseWriter, r *http.Request) {

	var loginData LoginData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)

	}

	log.Printf("login: %+v", loginData.Form_email)
	log.Printf("pass: %+v", loginData.Form_password)

	// Здесь вы можете обработать данные (например, проверить пароль) и отправить ответ на фронтенд

}

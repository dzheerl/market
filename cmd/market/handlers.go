package main

import (
	"encoding/json"
	// "database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

}



func sendUser(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	var request LoginData

	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

    query := "SELECT uid FROM user WHERE uid = ?"
    var value string
    err = db.QueryRow(query, request.Form_email).Scan(&value)
    if err != nil {
        fmt.Println(err)
		log.Println("test3")
		value = ""
        // return
    }
	// log.Println("This is ", value);
	if value != "" {
		fmt.Fprintf(w, "exists")
		log.Println("test2")
	} else {
		err = createUser(db, request)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		} else {
			fmt.Fprintf(w, "Succes")
		}
	}
	// if value == "" {
	// 	err = createUser(db, request)
	// 	if err != nil {
	// 		http.Error(w, "Internal Server Error", 500)
	// 		log.Println(err.Error())
	// 		return
	// 	}
	// } 
}
}

func createUser(db *sqlx.DB, request LoginData) error {
	log.Println("test")
	const query = `
		INSERT INTO
			user
		(
			uid, password
		)
		VALUES
		(
			?,
			?
		)
	`
	log.Println("Succes")
    _, err := db.Exec(query, request.Form_email, request.Form_password) // Сами данные передаются через аргументы к ф-ии Exec
   return err
}

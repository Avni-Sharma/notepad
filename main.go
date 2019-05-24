package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func main() {
	dbdetails := "host=localhost port=5430 user=postgres dbname=postgres password=secret sslmode=disable"
	DB, err := sql.Open("postgres", dbdetails)
	if err != nil {
		log.Fatal(err)
	}
	DB.Exec("SELECT 1")

	http.HandleFunc("/note", createNoteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	if r.Method == "POST" {

		body, _ := ioutil.ReadAll(r.Body)
		bodyText := string(body)
		//insert the body into the database
		DB.Exec(`INSERT INTO "note" (note) VALUES ($1)`, bodyText)
		return
	}
	if r.Method == "GET" {
		//read from db and print
	}

}

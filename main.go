package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func main() {
	var err error
	dbdetails := "host=localhost port=5430 user=postgres dbname=postgres password=secret sslmode=disable"
	DB, err = sql.Open("postgres", dbdetails)
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
		sqlStatement := `INSERT INTO note (note) VALUES ($1)`
		_, err1 := DB.Exec(sqlStatement, bodyText)
		if err1 != nil {
			fmt.Println("Error-----------", err1)
		}
		return
	}
	if r.Method == "GET" {
		sqlStatement := `SELECT note FROM note`
		rows, err1 := DB.Query(sqlStatement)
		if err1 != nil {
			fmt.Println("Error-----------", err1)
		}
		var objs []string
		for rows.Next() {
			s := ""
			err2 := rows.Scan(&s)
			if err2 != nil {
				fmt.Print(err2)
			}
			objs = append(objs, s)
		}
		result := strings.Join(objs, "\n")
		fmt.Fprintf(w, result)
		return
	}

}

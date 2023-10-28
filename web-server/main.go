package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type user struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

func main() {
	db, dbErr := sql.Open("mysql", "root:password@(127.0.0.1:3306)/test?parseTime=true")
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	_, createTableErr := db.Exec(query)

	if createTableErr != nil {
		log.Fatal(createTableErr)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!")
		fmt.Println(r.Method)
		fmt.Println(r.Host)
		fmt.Println(r.URL)
	})

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT * FROM users`)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()
		var users []user
		for rows.Next() {
			var u user
			rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			users = append(users, u)
		}
	})

	r.HandleFunc("/user/add/{username}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		password := "password"
		createdAt := time.Now()
		result, err := db.Exec(
			`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`,
			username, password, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		id, _ := result.LastInsertId()

		fmt.Fprintf(w, "User ID: %s", id)
	})

	r.HandleFunc("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["userId"]
		fmt.Fprintf(w, "Welcome user %s", userId)
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Looking at book %s at page %s", vars["title"], vars["page"])
	}).Methods("GET")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}

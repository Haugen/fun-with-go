package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!")
		fmt.Println(r.Method)
		fmt.Println(r.Host)
		fmt.Println(r.URL)
	})

	r.HandleFunc("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["userId"]
		fmt.Fprintf(w, "Welcome user %s", userId)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}

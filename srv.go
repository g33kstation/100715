package main

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	log.Println("Server started on :4400...")
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":4400", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	createTable()
	insertPerson()
	fmt.Fprint(w, "Helo")
}

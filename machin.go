package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Helo")
}

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":4400", nil)
}

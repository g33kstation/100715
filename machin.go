package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log" // logging operations
	"os" // os operations
)

func boing(){
	// remove db file if exists
	os.Remove("./test.db")

	// get connection to db
	db, err := sql.Open("sqlite3", "./test.db")
	// log error if any
	if err != nil {
		log.Fatal(err)
	}
	// ensure db is closed
	defer db.Close()

	sqlStmt := `
	create table if not exists person (id integer not null primary key, name text);
	`
	// log error if any
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Helo")
	boing()
}

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":4400", nil)
}

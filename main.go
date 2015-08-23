package main

import (
	"database/sql"
	"fmt"
	"log" // logging operations
	"os"  // os operations

	_ "github.com/mattn/go-sqlite3"
)

// La fonction init() est toujours appelé avant la fonction main()
//
// La différence avec `main` est qu'il ne peut exister qu'UNE fonction `main`,
// Mais qu'on peut mettre autant de function `init` qu'on veut ;)
//
// Utilisé idéalement pour toutes les actions à "pré-exécuter"
// avant le démarrage de notre programme
func init() {
	// Si une BDD existe, on la supprime avant de commencer
	if x, _ := os.Lstat("./test.db"); x != nil {
		os.Remove("./test.db")
	}
}

func createTable() {

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

func insertPerson() {
	// get connection to db
	db, err := sql.Open("sqlite3", "./test.db")
	// log error if any
	if err != nil {
		log.Fatal(err)
	}
	// ensure db is closed
	defer db.Close()

	// check if db alive
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// simple datas insertion
	_, err = db.Exec("insert into person(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	// transaction based data insertion

	// prepare a transaction object
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// prepare the statement
	stmt, err := tx.Prepare("insert into person(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// for loop insertion
	for k, v := range map[int]string{5: "pif", 6: "pof", 7: "paf"} {
		_, err = stmt.Exec(k, fmt.Sprintf(v))
		if err != nil {
			log.Fatal(err)
		}
	}
	// commit transaction
	tx.Commit()

}

func main() {
	runServer()
}

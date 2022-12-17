package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	id       int
	name     string
	nickname string
)

func main() {
	database, _ := sql.Open("sqlite3", "alem.db")
	// fmt.Println(reflect.TypeOf(database))

	Check(database)
	Add(database)
	// Delete(database)
	Select(database)
}

func Check(db *sql.DB) {
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, name TEXT, nickname TEXT)")
	statement.Exec()
}

func Add(db *sql.DB) {
	statement, _ := db.Prepare("INSERT INTO people (name, nickname) VALUES (?, ?)")
	statement.Exec("tamer", "quejey")
}

func Delete(db *sql.DB) {
	for i := 3; i < 9; i++ {
		res, err := db.Exec("delete from people where id = $1", i)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.RowsAffected())
	}
}

func Select(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, nickname FROM people")
	if err != nil {
		fmt.Println("error: select")
	}

	for rows.Next() {
		rows.Scan(&id, &name, &nickname)
		fmt.Printf("%d) %s %s\n", id, name, nickname)
	}
}

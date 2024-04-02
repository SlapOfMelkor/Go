package main

import (
	"database/sql"
	"log"

	"Apigodev/router"
	"Apigodev/tasks"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgresql://username:password@postgres/todo_list?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tasks.InitDB(db)

	router.StartServer()
}

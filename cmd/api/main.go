package main

import (
	"database/sql"
	"log"
)

func main() {
	// create database connection
	db, err := sql.Open("sqlite3","./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


}
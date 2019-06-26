package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func getConnection() *sql.DB {

	dsn := "postgres://luis.farfan.lara:password@127.0.0.1:5432/music?sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

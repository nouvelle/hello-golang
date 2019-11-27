package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Gopher struct {
	Id    int
	Name  string
	Img   string
	Count int
}

// 1. connect to Database
func ConnectDB() (db *sql.DB) {
	db, err := sql.Open("postgres", "user=eriko dbname=golang sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db

	//INSERT
	// var gopherID string
	// id := 1
	// name := "coffee"
	// img := "main.png"
	// count := 0
	// err = db.QueryRow("INSERT INTO gopher(id, name, img, count) VALUES($1,$2,$3,$4) RETURNING id", id, name, img, count).Scan(&gopherID)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(gopherID)
}

package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Menu struct {
	Id    int
	Name  string
	Price int
}

// 1. connect to Database
func ConnectDB() (db *sql.DB) {
	db, err := sql.Open("postgres", "user=eriko dbname=golang sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db

	//INSERT
	// var menuID string
	// id := 1
	// menu := "coffee"
	// price := 350
	// err = Db.QueryRow("INSERT INTO menu(id, manu, price) VALUES($1,$2,$3) RETURNING id", id, menu, price).Scan(&menuID)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(menuID)
}

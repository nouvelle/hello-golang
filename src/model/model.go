package model

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

var Db *sql.DB

// 1. connect to Database
func ConnectDB() {
	var err error
	Db, err = sql.Open("postgres", "user=eriko dbname=golang sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//INSERT
	var menuID string
	id := 1
	menu := "coffee"
	price := 350
	err = Db.QueryRow("INSERT INTO menu(id, manu, price) VALUES($1,$2,$3) RETURNING id", id, menu, price).Scan(&menuID)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(menuID)

}

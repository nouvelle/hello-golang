package model

import (
	"database/sql"
	"fmt"

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
}

// 2.SELECT ALL DATA
func GetALL() []Gopher {
	gophers := []Gopher{}
	db := ConnectDB()

	result, err := db.Query("SELECT * FROM gopher ORDER BY id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		gopher := Gopher{}
		var id int
		var name string
		var img string
		var count int

		err = result.Scan(&id, &name, &img, &count)
		if err != nil {
			panic(err.Error())
		}
		gopher.Id = id
		gopher.Name = name
		gopher.Img = img
		gopher.Count = count
		gophers = append(gophers, gopher)
	}
	defer db.Close()

	fmt.Println(gophers)
	return gophers
}

// 3.Update Gopher's like count
func PostCount(reqCnt int, reqId int) {
	db := ConnectDB()

	err := db.QueryRow("UPDATE gopher SET count = $1 WHERE id = $2", reqCnt, reqId)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

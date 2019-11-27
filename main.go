package main

import (
	"database/sql"
	"fmt"

	"github.com/eriko/to-do/src/controller"

	"github.com/gin-gonic/gin"

	// postgres driver
	_ "github.com/lib/pq"
)

// type Menu struct {
// 	Id    int
// 	Name  string
// 	Price int
// }

func main() {
	router := gin.Default()

	// connect database
	db, err := sql.Open("postgres", "user=eriko dbname=golang sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// INSERT
	var lastInsertId int
	err = db.QueryRow("INSERT INTO menu(name,price) VALUES($1,$2) returning id;", "coffee", 350).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertId)

	// SELECT ALL DATA
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM menu")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var price int
		err = rows.Scan(&id, &name, &price)
		if err != nil {
			panic(err)
		}
		fmt.Println("id | name | price")
		fmt.Printf("%3v | %8v  | %6v\n", id, name, price)
	}

	// API namespace
	// v1 := router.Group("/api/v1")
	// {
	// 	v1.GET("/menus", controller.GetMenu)
	// }

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}

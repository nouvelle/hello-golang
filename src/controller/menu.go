package controller

import (
	"fmt"
	"net/http"

	"github.com/eriko/to-do/src/model"
	"github.com/gin-gonic/gin"
)

// 2. get data
func GetAllMenu(c *gin.Context) {
	db := model.ConnectDB()
	// // err = Db.QueryRow("SELECT id, menu, price FROM menu where id = $1", id).Scan(&menu.Id, &menu.Name, &menu.Price)
	// menu, err := Db.Query("SELECT * FROM menu)
	// return menu

	// SELECT ALL DATA
	result, err := db.Query("SELECT * FROM menu")
	if err != nil {
		panic(err.Error())
	}

	menus := []model.Menu{}

	for result.Next() {
		menu := model.Menu{}
		var id int
		var name string
		var price int

		err = result.Scan(&id, &name, &price)
		if err != nil {
			panic(err.Error())
		}
		menu.Id = id
		menu.Name = name
		menu.Price = price
		menus = append(menus, menu)
		// fmt.Println("id | name | price")
		// fmt.Printf("%3v | %8v  | %6v\n", id, name, price)
	}
	fmt.Println(menus)
	c.JSON(http.StatusOK, gin.H{"menus": menus})
	// result.Close()
	// return
}

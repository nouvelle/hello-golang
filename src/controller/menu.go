package controller

import (
	"fmt"
	"net/http"

	"github.com/eriko/to-do/src/model"
	"github.com/gin-gonic/gin"
)

// 2. get data
func GetAllGopher(c *gin.Context) {
	db := model.ConnectDB()

	// SELECT ALL DATA
	result, err := db.Query("SELECT * FROM gopher")
	if err != nil {
		panic(err.Error())
	}

	gophers := []model.Gopher{}

	for result.Next() {
		gopher := model.Gopher{}
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
	fmt.Println(gophers)
	c.JSON(http.StatusOK, gin.H{"gophers": gophers})
}

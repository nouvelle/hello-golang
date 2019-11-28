package controller

import (
	"fmt"
	"net/http"

	"github.com/eriko/to-do/src/model"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

// get data
func GetAllGopher(c *gin.Context) {
	db := model.ConnectDB()

	// SELECT ALL DATA
	result, err := db.Query("SELECT * FROM gopher ORDER BY id")
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

// Update Gopher's like count
func UpdataGopher(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqId := post.Id
	reqCnt := post.Count
	// fmt.Println("reqId: %s; reqCnt: %s", reqId, reqCnt)

	// SELECT ALL DATA
	db := model.ConnectDB()

	err := db.QueryRow("UPDATE gopher SET count = $1 WHERE id = $2", reqCnt, reqId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success count up"})
}

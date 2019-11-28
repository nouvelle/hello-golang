package controller

import (
	"net/http"

	db "github.com/eriko/to-do/src/model"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

// GET ALL DATA
func GetAllGopher(c *gin.Context) {
	resultAllGophers := db.GetALL()
	c.JSON(http.StatusOK, gin.H{"gophers": resultAllGophers})
}

// Update Gopher's like count
func UpdataGopher(c *gin.Context) {
	post := Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqId := post.Id
	reqCnt := post.Count

	db.PostCount(reqCnt, reqId)
	c.JSON(http.StatusOK, gin.H{"status": "success count up"})
}

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

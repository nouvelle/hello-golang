package main

import (
	"net/http"

	"github.com/eriko/gophers/src/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// start up serve
	serve()
}

func serve() {
	router := gin.Default()

	// for CORS
	config := cors.DefaultConfig()
	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	// router.StaticFS("/gophers", http.Dir("../dist"))
	router.StaticFS("/gogo", http.Dir("./dist"))

	config.AllowOrigins = []string{"https://go-go-golang.herokuapp.com", "http://localhost:8081"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	router.Use(cors.New(config))

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/gophers", controller.GetAllGopher)
		v1.POST("/gophers", controller.UpdataGopher)
	}

	// router.GET("/", controller.IndexGET)

	// router.GET("/someGet", getting)
	// router.PATCH("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	router.Run(":8081")
}

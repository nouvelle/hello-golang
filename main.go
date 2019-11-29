package main

import (
	"net/http"
	"os"

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

	// for router
	router.StaticFS("/gogo", http.Dir("./dist"))

	config.AllowOrigins = []string{"https://go-go-golang.herokuapp.com", "http://localhost:8081"}
	router.Use(cors.New(config))

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/gophers", controller.GetAllGopher)
		v1.POST("/gophers", controller.UpdataGopher)
	}

	// router.GET("/", controller.IndexGET)
	router.Run(":8081")

	router.Run(":" + os.Getenv("PORT"))
}

package main

import (
	"github.com/eriko/to-do/src/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// for CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	router.Use(cors.New(config))

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/gophers", controller.GetAllGopher)
	}

	router.GET("/", controller.IndexGET)

	// router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	router.Run(":8081")
}

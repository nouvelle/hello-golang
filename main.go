package main

import (
	"github.com/eriko/to-do/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/menus", controller.GetAllMenu)
	}

	router.GET("/", controller.IndexGET)

	// router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	router.Run(":8080")
}

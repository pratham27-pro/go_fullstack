package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pratham27-pro/go_fullstack/server/routes"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

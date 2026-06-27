package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pratham27-pro/go_fullstack/server/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movies/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie ", controllers.AddMovie())

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

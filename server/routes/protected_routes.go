package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pratham27-pro/go_fullstack/server/controllers"
	"github.com/pratham27-pro/go_fullstack/server/middlware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middlware.AuthMiddleware())

	router.GET("/movies/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())
}

package main

import (
	"gin-app/config"
	"gin-app/middlewares"
	"gin-app/repositories"
	"gin-app/routes"
	"gin-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	db := config.ConnectDatabase()

	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)

	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middlewares.ErrorMiddleware())

	routes.UserRoutes(r, service)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9000")
}

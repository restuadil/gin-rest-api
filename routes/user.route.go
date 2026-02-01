package routes

import (
	"gin-app/controllers"
	"gin-app/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, service *services.UserService) {
	r.GET("/users", controllers.GetUsers(service))
	r.GET("/users/:id", controllers.GetUserByID(service))
}

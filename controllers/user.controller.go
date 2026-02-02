package controllers

import (
	"fmt"
	"gin-app/dto"
	"gin-app/errors"
	"gin-app/responses"
	"gin-app/services"
	"gin-app/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(service *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := service.GetUsers(c.Request.Context())
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, responses.Success(users, "Users retrieved successfully", http.StatusOK))
	}
}

func GetUserByID(service *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		var id int64
		_, err := fmt.Sscan(idParam, &id)

		if err != nil {
			c.Error(errors.NewValidationError("Invalid Id Format"))
			return
		}
		user, err := service.GetUserByID(c.Request.Context(), id)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, responses.Success(user, "User retrieved successfully", http.StatusOK))
	}
}

func CreateUser(service *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var createUserDTO dto.CreateUserDTO

		if err := c.ShouldBindJSON(&createUserDTO); err != nil {
			c.Error(errors.NewValidationError("Invalid Request Body"))
			return
		}

		if err := validation.Validate.Struct(createUserDTO); err != nil {
			message := validation.FormatValidationError(err)
			c.Error(errors.NewValidationError(message))
			return
		}

		user, err := service.CreateUser(c.Request.Context(), createUserDTO)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusCreated, responses.Success(responses.ToUserResponse(user), "User created successfully", http.StatusCreated))
	}
}

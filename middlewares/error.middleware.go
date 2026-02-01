package middlewares

import (
	"gin-app/errors"
	"gin-app/responses"

	"github.com/gin-gonic/gin"
)


func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		switch e := err.(type) {
		case *errors.AppError:
			c.JSON(e.Code, responses.Error(e.Code, e.Type,e.Message))
		default:
			c.JSON(500, responses.Error(500,"Internal Server Error", "Something went wrong"))
		}
	}
}

package responses

import "github.com/gin-gonic/gin"

func Success(data interface{}, message string, code int) gin.H {
	return gin.H{
		"status": true,
		"statusCode":  code,
		"message": message,
		"data":    data,
	}
}

func Error(code int,err string, message string) gin.H {
	return gin.H{
		"statusCode":  code,
		"status": false,
		"error":   err,
		"message": message,
	}
}

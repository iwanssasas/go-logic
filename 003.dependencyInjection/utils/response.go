package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func Response(data interface{}) gin.H {
	return gin.H{
		"data":  data,
		"error": nil,
	}
}

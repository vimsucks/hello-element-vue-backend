package model

import "github.com/gin-gonic/gin"

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"data":    data,
		"success": true,
	}
}

func FailureResponse(code int, reason string) gin.H {
	return gin.H{
		"success": false,
		"code":    code,
		"reason":  reason,
	}
}

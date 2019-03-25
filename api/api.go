package api

import "github.com/gin-gonic/gin"

func RegisterRouteGroup(router *gin.RouterGroup) {
	router.GET("/guitar/score/all", getAllScore)
	router.POST("/guitar/score", createNewScore)
	router.DELETE("/guitar/score/:uid", deleteScore)
}
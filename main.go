package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-element-vue-backend/api"
	"hello-element-vue-backend/config"
	"hello-element-vue-backend/model"
	"log"
	"net/http"
	"github.com/gin-contrib/cors"
)

func init() {
	config.InitFlagAndConfig()
	model.InitDatabase()
}

func main() {
	router := gin.Default()
	// Default Allow All Origins
	router.Use(cors.Default())

	apiRouter := router.Group("/api")
	api.RegisterRouteGroup(apiRouter)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	err := router.Run(fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port))
	if err != nil {
		log.Fatal(err)
	}
}

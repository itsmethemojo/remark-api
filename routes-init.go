package main

import (
	"docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

var (
	router = gin.Default()
)

// @title reMARK
// @version 1.0
// @description an API to bookmark URLs and also sort them over time by usage or re-bookmarks
// @BasePath /v1
func RoutesRun() {
	docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, os.Getenv("SCHEMA"))
	docs.SwaggerInfo.Host = os.Getenv("HOST") + ":" + os.Getenv("PORT")
	url := ginSwagger.URL(os.Getenv("SCHEMA") + "://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			//TODO answer in json
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	//TODO maybe use this to also host frontend
	// router.Static("/static", "./static")
	getRoutes()
	router.Run(":" + os.Getenv("PORT"))
}

func getRoutes() {
	v1 := router.Group("/v1")
	addBookmarkRoutes(v1)
}

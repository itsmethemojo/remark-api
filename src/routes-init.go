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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN"))
		c.Header("Access-Control-Allow-Credentials", os.Getenv("ACCESS_CONTROL_ALLOW_CREDENTIALS"))
		c.Header("Access-Control-Allow-Headers", os.Getenv("ACCESS_CONTROL_ALLOW_HEADERS"))
		c.Header("Access-Control-Allow-Methods", os.Getenv("ACCESS_CONTROL_ALLOW_METHODS"))
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// @title reMARK
// @version 1.0
// @description an API to bookmark URLs and also sort them over time by usage or re-bookmarks
// @BasePath /v1
func RoutesRun() {
	if os.Getenv("CORS_ENABLED") == "1" {
		router.Use(CORSMiddleware())
	}

	host_with_port := os.Getenv("APP_DOMAIN") + ":" + os.Getenv("APP_PORT")
	schema := os.Getenv("APP_SCHEMA")
	swagger_path := os.Getenv("SWAGGER_PATH")
	base_path := os.Getenv("API_PATH_PREFIX") + "/v1"

	docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, schema)
	docs.SwaggerInfo.Host = host_with_port
	docs.SwaggerInfo.BasePath = base_path

	url := ginSwagger.URL(schema + "://" + host_with_port + swagger_path + "/doc.json")
	router.GET(swagger_path+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			//TODO answer in json
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	//define routes

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	v1 := router.Group(base_path)
	addBookmarkRoutes(v1)

	router.LoadHTMLGlob("templates/*")

	if os.Getenv("LOGIN_PROVIDER") == "DEX" {
		auth := router.Group("auth")
		addAuthRoutes(auth)
	}

	err := router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic("starting webserver failed")
	}
}

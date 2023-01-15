package main

import (
	"docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var (
	router = gin.Default()
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", (EnvHelper).Get(EnvHelper{}, "ACCESS_CONTROL_ALLOW_ORIGIN"))
		c.Header("Access-Control-Allow-Credentials", (EnvHelper).Get(EnvHelper{}, "ACCESS_CONTROL_ALLOW_CREDENTIALS"))
		c.Header("Access-Control-Allow-Headers", (EnvHelper).Get(EnvHelper{}, "ACCESS_CONTROL_ALLOW_HEADERS"))
		c.Header("Access-Control-Allow-Methods", (EnvHelper).Get(EnvHelper{}, "ACCESS_CONTROL_ALLOW_METHODS"))
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
	if (EnvHelper).Get(EnvHelper{}, "CORS_ENABLED") == "1" {
		router.Use(CORSMiddleware())
	}

	host_with_port := (EnvHelper).Get(EnvHelper{}, "SWAGGER_HOST") + ":" + (EnvHelper).Get(EnvHelper{}, "SWAGGER_PORT")
	swagger_schema := (EnvHelper).Get(EnvHelper{}, "SWAGGER_SCHEMA")
	swagger_path := (EnvHelper).Get(EnvHelper{}, "SWAGGER_PATH")
	base_path := (EnvHelper).Get(EnvHelper{}, "API_PATH_PREFIX") + "/v1"

	docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, swagger_schema)
	docs.SwaggerInfo.Host = host_with_port
	docs.SwaggerInfo.BasePath = base_path

	url := ginSwagger.URL(swagger_schema + "://" + host_with_port + swagger_path + "/doc.json")
	router.GET(swagger_path+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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

	//define routes
	v1 := router.Group(base_path)
	addBookmarkRoutes(v1)

	if (EnvHelper).Get(EnvHelper{}, "LOGIN_PROVIDER") == "DEX" {
		router.LoadHTMLGlob("templates/*")
		auth := router.Group("auth")
		addAuthRoutes(auth)
	}

	err := router.Run(":" + (EnvHelper).Get(EnvHelper{}, "PORT"))
	if err != nil {
		panic("starting webserver failed")
	}
}

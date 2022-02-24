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

	docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, (EnvHelper).Get(EnvHelper{}, "SWAGGER_SCHEMA"))
	docs.SwaggerInfo.Host = (EnvHelper).Get(EnvHelper{}, "SWAGGER_HOST") + ":" + (EnvHelper).Get(EnvHelper{}, "SWAGGER_PORT")
	url := ginSwagger.URL((EnvHelper).Get(EnvHelper{}, "SWAGGER_SCHEMA") + "://" + (EnvHelper).Get(EnvHelper{}, "SWAGGER_HOST") + ":" + (EnvHelper).Get(EnvHelper{}, "SWAGGER_PORT") + (EnvHelper).Get(EnvHelper{}, "SWAGGER_PATH") + "/doc.json") //TODO make swagger path to env var
	router.GET((EnvHelper).Get(EnvHelper{}, "SWAGGER_PATH") + "/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url)) //TODO make swagger path to env var
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
	router.Run(":" + (EnvHelper).Get(EnvHelper{}, "PORT"))
}

func getRoutes() {
	v1 := router.Group((EnvHelper).Get(EnvHelper{}, "API_PATH_PREFIX") + "/v1") //TODO make swagger pre path before /v1 to env var
	addBookmarkRoutes(v1)
}

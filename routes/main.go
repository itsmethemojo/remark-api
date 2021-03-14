package routes

import (
	"../docs"
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
func Run() {
	//TODO rename this .env.default variable
	docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, os.Getenv("HTTPS_OR_HTTPS"))
	docs.SwaggerInfo.Host = os.Getenv("HOST") + ":" + os.Getenv("PORT")
	//TODO read host,schemes from .env with useful defaults
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
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

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	addBookmarkRoutes(v1)

	//v2 := router.Group("/v2")
	//addBookmarkRoutes(v2)
}

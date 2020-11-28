package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../models/bookmark"
)

func addBookmarkRoutes(rg *gin.RouterGroup) {

	bookmarks := rg.Group("/bookmark")

	bookmarks.GET("/", func(c *gin.Context) {
		// TODO handle authentification later
		b := bookmark.BookmarkModel{UserId: 1}
		return_data, err := b.GetAll()
		if err == nil {
			c.JSON(http.StatusOK, return_data)
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]string{"message": "ohoh"})
		}
	})
}

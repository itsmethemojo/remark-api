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
		b := bookmark.BookmarkModel{}
		return_data, err := b.ListAll(c.Query("user_id"))
		if err == nil {
			c.JSON(http.StatusOK, return_data)
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
			//c.JSON(http.StatusInternalServerError, map[string]string{"message": userId})
		}
	})

	bookmarks.GET("/remark/", func(c *gin.Context) {
		// TODO handle authentification later
		b := bookmark.BookmarkModel{}
		remarkErr := b.Remark(c.Query("user_id"), c.Query("remark"))
		if remarkErr == nil {
			c.JSON(http.StatusOK, "everything is fine. TODO add status array")
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
			//c.JSON(http.StatusInternalServerError, map[string]string{"message": userId})
		}
	})
}

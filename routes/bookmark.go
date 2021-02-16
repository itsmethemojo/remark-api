package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"../models/bookmark"
)

func addBookmarkRoutes(rg *gin.RouterGroup) {

	bookmarks := rg.Group("/bookmark")

	bookmarks.GET("/", func(c *gin.Context) {
		// TODO handle authentification later
		userId, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
		if err != nil {
			// do something
		}
		b := bookmark.BookmarkModel{UserId: userId}
		return_data, err := b.ListAll()
		if err == nil {
			c.JSON(http.StatusOK, return_data)
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]uint64{"message": userId})
			//c.JSON(http.StatusInternalServerError, map[string]string{"message": userId})
		}
	})

	bookmarks.GET("/remark/", func(c *gin.Context) {
		// TODO handle authentification later
		userId, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
		if err != nil {
			// do something
		}
		b := bookmark.BookmarkModel{UserId: userId}
		remarkErr := b.Remark(c.Query("remark"))
		if remarkErr == nil {
			c.JSON(http.StatusOK, "everything is fine. TODO add status array")
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]uint64{"message": userId})
			//c.JSON(http.StatusInternalServerError, map[string]string{"message": userId})
		}
	})
}

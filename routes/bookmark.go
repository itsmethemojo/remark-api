package routes

import (
	"net/http"
	//"fmt"
	"log"
	"github.com/gin-gonic/gin"

	. "../models/bookmark"
	//repository "../repositories/bookmark"
)

func addBookmarkRoutes(rg *gin.RouterGroup) {

	bookmarks := rg.Group("/bookmark")

	bookmarks.GET("/", func(c *gin.Context) {
		// TODO handle authentification later
		// right now user_id parameter is optional and can be extracted by validating the cookie
		// if user_id gets mandatory it should not be a query but a body parameter
		// maybe other parameters should move to body to
		// TODO add demo mode to skip authentification
		b := BookmarkModel{}
		return_data, err := b.ListAll(c.Query("user_id"))
		//allBookmarkData := repository.AllBookmarkData{
		//	Bookmarks: return_data.Bookmarks,
		//}

		log.Println("%v", return_data)
		if err == nil {
			c.JSON(http.StatusOK, return_data)
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
			//c.JSON(http.StatusInternalServerError, map[string]string{"message": userID})
		}
	})

	bookmarks.GET("/remark/", func(c *gin.Context) {
		// TODO handle authentification later
		b := BookmarkModel{}
		remarkError := b.Remark(c.Query("user_id"), c.Query("remark"))
		if remarkError == nil {
			c.JSON(http.StatusOK, "everything is fine. TODO add status array")
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
			//c.JSON(http.StatusInternalServerError, map[string]string{"message": userID})
		}
	})

	bookmarks.GET("/click/", func(c *gin.Context) {
		// TODO handle authentification later
		b := BookmarkModel{}
		clickError := b.Click(c.Query("user_id"), c.Query("id"))
		if clickError == nil {
			c.JSON(http.StatusOK, "everything is fine. TODO add status array")
		} else {
			// TODO retreive response text and code from model -> no error handling needed
			c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
			//c.JSON(http.StatusInternalServerError, map[string]string{"message": userID})
		}
	})
}

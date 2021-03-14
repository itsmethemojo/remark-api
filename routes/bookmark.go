package routes

import (
	"net/http"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"log"

	. "../models/authentification"
	. "../models/bookmark"
	//repository "../repositories/bookmark"
)

func addBookmarkRoutes(rg *gin.RouterGroup) {

	bookmarks := rg.Group("/bookmark")
	bookmarks.GET("/", routeBookmarks)
	bookmarks.GET("/remark/", routeBookmarksRemark)
	bookmarks.GET("/click/", routeBookmarksClick)
}

// @Description get all bookmarks for user
// @ID bookmark
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param AUTH_TOKEN header string true "authorization token"
// @router /bookmark/ [get]
func routeBookmarks(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("AUTH_TOKEN"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{}
	return_data, err := b.ListAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
		return
	}
	c.JSON(http.StatusOK, return_data)
}

//TODO seitch to post

// @Description bookmark an url
// @ID bookmark-remark
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param AUTH_TOKEN header string true "authorization token"
// @Param remark query string true "url to be bookmarked"
// @router /bookmark/remark/ [get]
func routeBookmarksRemark(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("AUTH_TOKEN"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{}
	remarkError := b.Remark(userID, c.Query("remark"))
	if remarkError != nil {
		// TODO retreive response text and code from model -> no error handling needed
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"message": "ok"})
}

//TODO seitch to post

// @Description save a bookmark click
// @ID bookmark-click
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param AUTH_TOKEN header string true "authorization token"
// @Param id query string true "bookmark id of the clicked bookmark"
// @router /bookmark/click/ [get]
func routeBookmarksClick(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("AUTH_TOKEN"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{}
	clickError := b.Click(userID, c.Query("id"))
	if clickError != nil {
		http_code := http.StatusInternalServerError
		message := "Internal Server Error"
		//TODO look up error message given to modifix http_code and message
		c.JSON(http_code, map[string]string{"message": message})
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"message": "ok"})
}

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type CreateJSONResult struct {
	Message string `json:"message"`
}

type AllDataJSONResult struct {
	Bookmarks []BookmarkEntity
	Remarks   []RemarkEntity
	Clicks    []ClickEntity
}

var database *gorm.DB

func addBookmarkRoutes(rg *gin.RouterGroup) {

	dsn := "host=" + os.Getenv("DATABASE_HOST") +
		" user=" + os.Getenv("DATABASE_USERNAME") +
		" password=" + os.Getenv("DATABASE_PASSWORD") +
		" dbname=" + os.Getenv("DATABASE_NAME") +
		" port=" + os.Getenv("DATABASE_PORT") +
		" sslmode=" + os.Getenv("DATABASE_SSLMODE") +
		" TimeZone=" + os.Getenv("DATABASE_TIMEZONE")
	var connectError error
	database, connectError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if connectError != nil {
		panic("could not connect to database")
	}

	bookmarks := rg.Group("/bookmark")
	bookmarks.GET("/", routeBookmarks)
	bookmarks.POST("/remark/", routeBookmarksRemark)
	bookmarks.POST("/click/", routeBookmarksClick)
	bookmarks.POST("/:id/", routeBookmarksEdit)
	bookmarks.DELETE("/:id/", routeBookmarksDelete)

	if os.Getenv("TEST_MODE") == "true" {
		bookmarks.DELETE("/", routeBookmarksDeleteAllData)
	}
}

// @Description get all bookmarks for user
// @ID bookmark
// @Accept json
// @Produce json
// @Success 200 {object} AllDataJSONResult{} "All Bookmark Data for User"
// @Param Authorization header string true "authorization token" default(LOCAL_TEST_TOKEN_1)
// @router /bookmark/ [get]
func routeBookmarks(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("Authorization"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{database}
	return_data, err := b.ListAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
		return
	}
	c.JSON(http.StatusOK, return_data)
}

//TODO this body definition just works for one parameter

// @Description bookmark an url
// @ID bookmark-remark
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 201 {object} CreateJSONResult{} "Entity inserted"
// @Param Authorization header string true "authorization token" default(LOCAL_TEST_TOKEN_1)
// @Param url body string true "url to be bookmarked, use format url="
// @router /bookmark/remark/ [post]
func routeBookmarksRemark(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("Authorization"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{database}
	modelError := b.Remark(userID, c.PostForm("url"))
	if modelError != nil {
		// TODO retreive response text and code from model -> no error handling needed
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "something is wrong"})
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"message": "ok"})
}

// @Description save a bookmark click
// @ID bookmark-click
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 201 {object} CreateJSONResult{} "Entity inserted"
// @Param Authorization header string true "authorization token" default(LOCAL_TEST_TOKEN_1)
// @Param id body string true "bookmark id of the clicked bookmark, use format id="
// @router /bookmark/click/ [post]
func routeBookmarksClick(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("Authorization"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{database}
	modelError := b.Click(userID, c.PostForm("id"))
	if modelError != nil {
		http_code := http.StatusInternalServerError
		message := "Internal Server Error"
		//TODO look up error message given to modifix http_code and message
		c.JSON(http_code, map[string]string{"message": message})
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"message": "ok"})
}

// @Description edit a bookmark
// @ID bookmark-edit
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 201 {object} CreateJSONResult{} "Entity updated"
// @Param Authorization header string true "authorization token" default(LOCAL_TEST_TOKEN_1)
// @Param id    path int    true "bookmark id"
// @Param title body string true "title to change, use format title="
// @router /bookmark/{id}/ [post]
func routeBookmarksEdit(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("Authorization"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{database}
	modelError := b.Edit(userID, c.Param("id"), c.PostForm("title"))
	if modelError != nil {
		http_code := http.StatusInternalServerError
		message := "Internal Server Error"
		//TODO look up error message given to modifix http_code and message
		c.JSON(http_code, map[string]string{"message": message})
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"message": "ok"})
}

// @Description delete a bookmark
// @ID bookmark-delete
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} CreateJSONResult{} "Entity deleted"
// @Param Authorization header string true "authorization token" default(LOCAL_TEST_TOKEN_1)
// @Param id    path int    true "bookmark id"
// @router /bookmark/{id}/ [delete]
func routeBookmarksDelete(c *gin.Context) {
	a := AuthentificationModel{}
	userID, authError := a.GetUserID(c.Request.Header.Get("Authorization"))
	if authError != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	b := BookmarkModel{database}
	modelError := b.Delete(userID, c.Param("id"))
	if modelError != nil {
		http_code := http.StatusInternalServerError
		message := "Internal Server Error"
		//TODO look up error message given to modifix http_code and message
		c.JSON(http_code, map[string]string{"message": message})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
func routeBookmarksDeleteAllData(c *gin.Context) {
	b := BookmarkModel{database}
	modelError := b.DeleteAllData()
	if modelError != nil {
		http_code := http.StatusInternalServerError
		message := "Internal Server Error"
		//TODO look up error message given to modifix http_code and message
		c.JSON(http_code, map[string]string{"message": message})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

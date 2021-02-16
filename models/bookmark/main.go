package bookmark

import (
	"strconv"
	//"errors"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	//. "../../entities/bookmark"
	bookmarkRepository "../../repositories/bookmark"
	//"time"
)

type BookmarkModel struct {
}

// @Description get all bookmarks for user
// @ID bookmark
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param user_id query int true "user id from bookmark owner"
// @router /bookmark/ [get]
func (this BookmarkModel) ListAll(userID string) (bookmarkRepository.AllBookmarkData, error) {
	// TODO also add http response code
	// right now error is always nil
	// TODO check authentification user should be logged in
	parsedUserId, parseErr := strconv.ParseUint(userID, 10, 32)
	if parseErr != nil {
		emptyData := bookmarkRepository.AllBookmarkData{}
		return emptyData, parseErr
	}
	bookmarkRepository := bookmarkRepository.BookmarkRepository{}
	return bookmarkRepository.ListAll(parsedUserId), nil //TODO check if err needed
}

// @Description bookmark an url
// @ID bookmark-remark
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param user_id query int true "user id from bookmark owner"
// @Param remark query string true "url to be bookmarked"
// @router /bookmark/remark/ [get]
func (this BookmarkModel) Remark(userID string, url string) error {
	parsedUserId, parseErr := strconv.ParseUint(userID, 10, 32)
	if parseErr != nil {
		return parseErr
	}
	bookmarkRepository := bookmarkRepository.BookmarkRepository{}
	repositoryError := bookmarkRepository.Remark(parsedUserId, url)
	return repositoryError
}

// @Description save a bookmark click
// @ID bookmark-click
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param user_id query int true "user id from bookmark owner"
// @Param id query string true "bookmark id of the clicked bookmark"
// @router /bookmark/click/ [get]
func (this BookmarkModel) Click(userID string, id string) error {
	parsedUserID, parsedUserIdError := strconv.ParseUint(userID, 10, 32)
	parsedID, parsedIDError := strconv.ParseUint(id, 10, 32)
	if parsedUserIdError != nil  {
		//TODO return message should say parameter type mismatch
		return parsedUserIdError
	}
	if parsedIDError != nil  {
		return parsedIDError
	}
	bookmarkRepository := bookmarkRepository.BookmarkRepository{}
	repositoryError := bookmarkRepository.Click(parsedUserID, parsedID)
	return repositoryError
}

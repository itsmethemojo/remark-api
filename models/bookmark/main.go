package bookmark

import (
	"strconv"
	//"errors"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	. "../../entities/bookmark"
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
func (b BookmarkModel) ListAll(userId string) ([]BookmarkEntity, error) {
	// TODO also add http response code
	// right now error is always nil
	// TODO check authentification user should be logged in
	parsedUserId, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		// do something
	}
	return bookmarkRepository.ListAll(parsedUserId), nil
}

// @Description bookmark an url
// @ID bookmark-remark
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param user_id query int true "user id from bookmark owner"
// @Param remark query string true "url to be bookmarked"
// @router /bookmark/remark/ [get]
func (b BookmarkModel) Remark(userId string, url string) error {
	parsedUserId, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		// do something
	}
	return bookmarkRepository.Remark(parsedUserId, url)
}

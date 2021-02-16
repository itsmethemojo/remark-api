package bookmark

import (
	//"errors"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	. "../../entities/bookmark"
	bookmarkRepository "../../repositories/bookmark"
	//"time"
)

type BookmarkModel struct {
	UserId uint64
}

// @Description get all bookmarks for user
// @ID bookmark
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param user_id query int true "user id from bookmark owner"
// @router /bookmark/ [get]
func (b BookmarkModel) ListAll() ([]BookmarkEntity, error) {
	// TODO also add http response code
	// right now error is always nil
	// TODO check authentification user should be logged in
	return bookmarkRepository.ListAll(b.UserId), nil
}

// @Description bookmark an url
// @ID bookmark-remark
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param user_id query int true "user id from bookmark owner"
// @Param remark query string true "url to be bookmarked"
// @router /bookmark/remark/ [get]
func (b BookmarkModel) Remark(url string) error {
	return bookmarkRepository.Remark(b.UserId, url)
}

package bookmark

import (
	"errors"
)

type BookmarkModel struct {
	UserId int
}

// @Description get Foo
// @ID get-foo
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /testapi/get-foo [get]mon error
// @router /bookmark [get]
func (b BookmarkModel) GetAll() (string, error) {
	// TODO also add http response code
	panic("foo")
	err_1 := errors.New("Error message_1: ")
	return "all bookmarks", err_1
}

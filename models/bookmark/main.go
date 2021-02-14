package bookmark

import (
	"errors"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"time"
	bookmarkRepository "../../repositories/bookmark"
)

type BookmarkModel struct {
	UserId int
}

type Product struct {
	ID        uint `gorm:"primaryKey"`
	Code      string
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
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
	//panic("foo")
	//bookmarkRepository.Remark("https://itsmethemojo.eu/remark/")
	bookmarkRepository.Remark("https://itsmethemojo.eu/")
	err_1 := errors.New("Error message_1: ")
	return "all bookmarks", err_1
}

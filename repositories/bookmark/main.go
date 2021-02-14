package bookmark

import (
	//"time"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "../../entities/bookmark"
)

//TODO extract database connect to a private function or init
func InitializeDatabase() {
	//TODO use https://github.com/joho/godotenv
	dsn := "root:rootpw@tcp(devdbhost:3306)/remark_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&BookmarkEntity{})
	if err != nil {
		panic("could not init database")
	}
}

func Remark(url string) {
	dsn := "root:rootpw@tcp(devdbhost:3306)/remark_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db.First(&product, 1)                 // find product with integer primary key
	var existingBookmarkEntity BookmarkEntity
	result := db.First(&existingBookmarkEntity, "url = ?", url) // find product with code D42

	//TODO https://github.com/dyatlov/go-htmlinfo
	// check for canonical url and title
	// if canonical url is given save that one?

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newBookmarkEntity := &BookmarkEntity{
			Url: url,
			UserId: 1, //TODO retrieve
			Title: url, //TODO retrieve bx curling url with same useragent
			RemarkCount: 1, //TODO check if it starts with 0
			ClickCount: 0,
		}
		db.Create(newBookmarkEntity)
		//TODO db.Create(newBookmarkRemarkEntity)
		return
	}

	existingBookmarkEntity.RemarkCount = existingBookmarkEntity.RemarkCount + 1
	db.Save(existingBookmarkEntity)
	//TODO only increase remark count and write remark entry
	// see https://gorm.io/docs/update.html
}

func click(bookmarkId int) {
	dsn := "root:rootpw@tcp(devdbhost:3306)/remark_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var existingBookmarkEntity BookmarkEntity
	result := db.First(&existingBookmarkEntity, bookmarkId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//TODO throw error
		return
	}

	//TODO only increase click count and write click entry


}

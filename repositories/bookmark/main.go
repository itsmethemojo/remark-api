package bookmark

import (
	//"time"
	. "../../entities/bookmark"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
//	"log"
)

type AllBookmarkData struct{
	Bookmarks  []BookmarkEntity
	Remarks    []RemarkEntity
	Clicks     []ClickEntity
}

type BookmarkRepository struct {
}

func (this BookmarkRepository) getDB() *gorm.DB {
	dsn := "root:rootpw@tcp(devdbhost:3306)/remark_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, connectError := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if connectError != nil {
		panic("failed to connect database")
	}
	return db
}

//TODO extract database connect to a private function or init
func (this BookmarkRepository) InitializeDatabase() {
	//TODO use https://github.com/joho/godotenv
	db := this.getDB()
	bookmarkEntityMigrateError := db.AutoMigrate(&BookmarkEntity{})
	remarkEntityMigrateError := db.AutoMigrate(&RemarkEntity{})
	clickEntityMigrateError := db.AutoMigrate(&ClickEntity{})
	if bookmarkEntityMigrateError != nil || remarkEntityMigrateError != nil || clickEntityMigrateError != nil  {
		panic("could not init database")
	}
}

func (this BookmarkRepository) ListAll(userID uint64) AllBookmarkData {
	db := this.getDB()
	var bookmarkEntities []BookmarkEntity
	db.Where("user_id = ?", userID).Find(&bookmarkEntities)
	var remarkEntities []RemarkEntity
	db.Raw("SELECT r.id, r.bookmark_id, r.created_at FROM bookmark_entities b JOIN remark_entities r ON b.id = r.bookmark_id WHERE b.user_id = ?", userID).Find(&remarkEntities)
	var clickEntities []ClickEntity
	db.Raw("SELECT c.id, c.bookmark_id, c.created_at FROM bookmark_entities b JOIN remark_entities c ON b.id = c.bookmark_id WHERE b.user_id = ?", userID).Find(&clickEntities)

	allBookmarkData := AllBookmarkData{
		Bookmarks: bookmarkEntities,
		Remarks: remarkEntities,
		Clicks: clickEntities,
	}
	return allBookmarkData
}

func (this BookmarkRepository) Remark(userID uint64, url string) error {
	db := this.getDB()
	//db.First(&product, 1)                 // find product with integer primary key
	var existingBookmarkEntity BookmarkEntity
	initialSearchResult := db.First(&existingBookmarkEntity, "url = ? AND user_id = ?", url, userID)
	//log.Println(result.RowsAffected)
	//log.Println(result.Error)
	//TODO https://github.com/dyatlov/go-htmlinfo
	// check for canonical url and title
	// if canonical url is given save that one?

	if errors.Is(initialSearchResult.Error, gorm.ErrRecordNotFound) {
		newBookmarkEntity := &BookmarkEntity{
			Url:         url,
			UserID:      userID, //TODO retrieve
			Title:       url,    //TODO retrieve bx curling url with same useragent
			RemarkCount: 1,      //TODO check if it starts with 0
			ClickCount:  0,
		}
		db.Create(newBookmarkEntity)
		searchResultAfterInsert := db.First(&existingBookmarkEntity, "url = ? AND user_id = ?", url, userID)
		if errors.Is(searchResultAfterInsert.Error, gorm.ErrRecordNotFound) {
			panic("this should never happen");
		}
	}

	newRemarkEntity := &RemarkEntity{
		BookmarkID:      existingBookmarkEntity.ID,
	}
	db.Create(newRemarkEntity)
	var bookmarkEntities []BookmarkEntity
	remarkCountResult := db.Raw("SELECT * FROM bookmark_entities b JOIN remark_entities r ON b.id = r.bookmark_id WHERE b.user_id = ?", userID).Find(&bookmarkEntities)
	existingBookmarkEntity.RemarkCount = uint64(remarkCountResult.RowsAffected)
	db.Save(existingBookmarkEntity)

	//TODO only increase remark count and write remark entry
	// see https://gorm.io/docs/update.html
	return nil
}

func (this BookmarkRepository) Click(userID uint64, bookmarkId uint64) error {
	db := this.getDB()
	var existingBookmarkEntity BookmarkEntity
	initialSearchResult := db.First(&existingBookmarkEntity, bookmarkId)
	if errors.Is(initialSearchResult.Error, gorm.ErrRecordNotFound) {
		panic("not found");
		// TODO return a not found error that causes notfound http respionse code
		return nil
	}
	newClickEntity := &ClickEntity{
		BookmarkID:      existingBookmarkEntity.ID,
	}
	db.Create(newClickEntity)
	var bookmarkEntities []BookmarkEntity
	remarkCountResult := db.Raw("SELECT * FROM bookmark_entities b JOIN click_entities c ON b.id = c.bookmark_id WHERE b.user_id = ?", userID).Find(&bookmarkEntities)
	existingBookmarkEntity.RemarkCount = uint64(remarkCountResult.RowsAffected)
	db.Save(existingBookmarkEntity)
	return nil
}

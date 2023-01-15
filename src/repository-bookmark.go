package main

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AllBookmarkData struct {
	Bookmarks []BookmarkEntity
	Remarks   []RemarkEntity
	Clicks    []ClickEntity
}

type BookmarkRepository struct {
}

func (this BookmarkRepository) getDB() (*gorm.DB, error) {
	dsn := "host=" + (EnvHelper).Get(EnvHelper{}, "DATABASE_HOST") +
		" user=" + (EnvHelper).Get(EnvHelper{}, "DATABASE_USERNAME") +
		" password=" + (EnvHelper).Get(EnvHelper{}, "DATABASE_PASSWORD") +
		" dbname=" + (EnvHelper).Get(EnvHelper{}, "DATABASE_NAME") +
		" port=" + (EnvHelper).Get(EnvHelper{}, "DATABASE_PORT") +
		" sslmode=" + (EnvHelper).Get(EnvHelper{}, "DATABASE_SSLMODE") +
		" TimeZone=" + (EnvHelper).Get(EnvHelper{}, "DATABASE_TIMEZONE")
	db, connectError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if connectError != nil {
		return db, errors.New("could not connect to database")
	}
	return db, nil
}

func (this BookmarkRepository) InitializeDatabase() error {
	db, dbConnectError := this.getDB()
	if dbConnectError != nil {
		return dbConnectError
	}
	bookmarkEntityMigrateError := db.AutoMigrate(&BookmarkEntity{})
	remarkEntityMigrateError := db.AutoMigrate(&RemarkEntity{})
	clickEntityMigrateError := db.AutoMigrate(&ClickEntity{})
	userEntityMigrateError := db.AutoMigrate(&UserEntity{})
	if bookmarkEntityMigrateError != nil || remarkEntityMigrateError != nil || clickEntityMigrateError != nil || userEntityMigrateError != nil {
		panic("could not init database")
	}
	return nil
}

func (this BookmarkRepository) DeleteAllData() error {
	db, dbConnectError := this.getDB()
	if dbConnectError != nil {
		return dbConnectError
	}

	db.Exec("DELETE FROM bookmark_entities")
	db.Exec("DELETE FROM remark_entities")
	db.Exec("DELETE FROM click_entities")
	db.Exec("DELETE FROM user_entities")

	return nil
}

func (this BookmarkRepository) ListAll(userID uint64) (AllBookmarkData, error) {
	db, dbConnectError := this.getDB()
	if dbConnectError != nil {
		return AllBookmarkData{}, dbConnectError
	}
	var bookmarkEntities []BookmarkEntity
	db.Where("user_id = ?", userID).Find(&bookmarkEntities)
	var remarkEntities []RemarkEntity
	db.Raw("SELECT r.id, r.bookmark_id, r.created_at FROM bookmark_entities b JOIN remark_entities r ON b.id = r.bookmark_id WHERE b.user_id = ? ORDER BY r.id DESC", userID).Find(&remarkEntities)
	var clickEntities []ClickEntity
	db.Raw("SELECT c.id, c.bookmark_id, c.created_at FROM bookmark_entities b JOIN click_entities c ON b.id = c.bookmark_id WHERE b.user_id = ? ORDER BY c.id DESC", userID).Find(&clickEntities)

	allBookmarkData := AllBookmarkData{
		Bookmarks: bookmarkEntities,
		Remarks:   remarkEntities,
		Clicks:    clickEntities,
	}
	return allBookmarkData, nil
}

func (this BookmarkRepository) Remark(userID uint64, url string) error {
	db, dbConnectError := this.getDB()
	if dbConnectError != nil {
		return dbConnectError
	}
	var existingBookmarkEntity BookmarkEntity
	initialSearchResult := db.First(&existingBookmarkEntity, "url = ? AND user_id = ?", url, userID)

	if errors.Is(initialSearchResult.Error, gorm.ErrRecordNotFound) {
		var title string
		title = url
		doc, fetchHtmlError := htmlquery.LoadURL(url)
		if fetchHtmlError == nil {
			titleNode := htmlquery.FindOne(doc, "//head/title")
			if titleNode != nil {
				title = htmlquery.InnerText(titleNode)
			}
		}
		newBookmarkEntity := &BookmarkEntity{
			Url:         url,
			UserID:      userID,
			Title:       title,
			RemarkCount: 1,
			ClickCount:  0,
		}
		db.Create(newBookmarkEntity)
		searchResultAfterInsert := db.First(&existingBookmarkEntity, "url = ? AND user_id = ?", url, userID)
		if errors.Is(searchResultAfterInsert.Error, gorm.ErrRecordNotFound) {
			panic("this should never happen")
		}
	}

	newRemarkEntity := &RemarkEntity{
		BookmarkID: existingBookmarkEntity.ID,
	}
	db.Create(newRemarkEntity)
	var bookmarkEntities []BookmarkEntity
	remarkCountResult := db.Raw("SELECT * FROM bookmark_entities b JOIN remark_entities r ON b.id = r.bookmark_id WHERE b.user_id = ? AND r.bookmark_id = ?", userID, existingBookmarkEntity.ID).Find(&bookmarkEntities)
	existingBookmarkEntity.RemarkCount = uint64(remarkCountResult.RowsAffected)
	db.Save(existingBookmarkEntity)
	return nil
}

func (this BookmarkRepository) Click(userID uint64, bookmarkId uint64) error {
	db, dbConnectError := this.getDB()
	if dbConnectError != nil {
		return dbConnectError
	}
	var existingBookmarkEntity BookmarkEntity
	initialSearchResult := db.First(&existingBookmarkEntity, bookmarkId)
	if errors.Is(initialSearchResult.Error, gorm.ErrRecordNotFound) {
		return errors.New("entity not found")
	}
	newClickEntity := &ClickEntity{
		BookmarkID: existingBookmarkEntity.ID,
	}
	db.Create(newClickEntity)
	var bookmarkEntities []BookmarkEntity
	clickCountResult := db.Raw("SELECT * FROM bookmark_entities b JOIN click_entities c ON b.id = c.bookmark_id WHERE b.user_id = ? AND c.bookmark_id = ?", userID, existingBookmarkEntity.ID).Find(&bookmarkEntities)
	existingBookmarkEntity.ClickCount = uint64(clickCountResult.RowsAffected)
	db.Save(existingBookmarkEntity)
	return nil
}

func (this BookmarkRepository) Edit(userID uint64, bookmarkId uint64, bookmarkTitle string) error {
	db, dbConnectError := this.getDB()
	if dbConnectError != nil {
		return dbConnectError
	}
	var existingBookmarkEntity BookmarkEntity
	initialSearchResult := db.First(&existingBookmarkEntity, "id = ? AND user_id = ?", bookmarkId, userID)

	if errors.Is(initialSearchResult.Error, gorm.ErrRecordNotFound) {
		return errors.New("entity not found")
	}

	existingBookmarkEntity.Title = bookmarkTitle
	db.Save(existingBookmarkEntity)
	return nil
}

func (this BookmarkRepository) Delete(userID uint64, bookmarkId uint64) error {
	db, dbConnectError := this.getDB()
	if dbConnectError != nil {
		return dbConnectError
	}
	var existingBookmarkEntity BookmarkEntity
	initialSearchResult := db.First(&existingBookmarkEntity, "id = ? AND user_id = ?", bookmarkId, userID)

	if errors.Is(initialSearchResult.Error, gorm.ErrRecordNotFound) {
		return errors.New("entity not found")
	}

	db.Exec("DELETE FROM click_entities WHERE bookmark_id = ?", bookmarkId)
	db.Exec("DELETE FROM remark_entities WHERE bookmark_id = ?", bookmarkId)
	db.Delete(existingBookmarkEntity)
	return nil
}

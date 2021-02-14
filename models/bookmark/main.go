package bookmark

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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

	dsn := "root:rootpw@tcp(devdbhost:3306)/remark_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

	err_1 := errors.New("Error message_1: ")
	return "all bookmarks", err_1
}

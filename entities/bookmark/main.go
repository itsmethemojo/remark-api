package bookmark

import (
	"time"
)

type BookmarkEntity struct {
	ID          uint `gorm:"primaryKey"`
	UserId      uint
	Url         string
	Title       string
	CustomTitle string
	RemarkCount uint
	ClickCount  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

package bookmark

import (
	"time"
)

type BookmarkEntity struct {
	ID          uint64 `gorm:"primaryKey"`
	UserId      uint64
	Url         string
	Title       string
	CustomTitle string
	RemarkCount uint64
	ClickCount  uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

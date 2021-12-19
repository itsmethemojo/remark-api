package main

import (
	"time"
)

type BookmarkEntity struct {
	ID          uint64 `gorm:"primaryKey"`
	UserID      uint64
	Url         string
	Title       string
	CustomTitle string
	RemarkCount uint64
	ClickCount  uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RemarkEntity struct {
	ID         uint64 `gorm:"primaryKey"`
	BookmarkID uint64
	CreatedAt  time.Time
}

type ClickEntity struct {
	ID         uint64 `gorm:"primaryKey"`
	BookmarkID uint64
	CreatedAt  time.Time
}

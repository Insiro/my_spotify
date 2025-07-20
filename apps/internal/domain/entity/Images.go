package entity

import (
	"time"
)

type Image struct {
	Id        uint      `gorm:"autoIncrement"`
	Url       string    `gorm:"size:500"`
	Height    int       ``
	Width     int       ``
	CreatedAt time.Time `gorm:"default:now()"`
}

package entity

import "time"

type PlayHistory struct {
	Id            uint          `gorm:"autoIncrement"`
	Track         Track         ``
	Album         Album         `gorm:"index:idx_play_history_album"`
	PrimaryArtist Artist        `gorm:"index:idx_play_history_primary_artist"`
	DurationMs    uint          `gorm:"default:0"`
	PlayedAt      time.Time     `gorm:"index:idx_play_history_played_at"`
	BlackListBy   []BlackListBy `gorm:"enum('artist')"`
	CreatedAt     time.Time     `gorm:"default:now()"`
}

const (
	BlackListByArtist BlackListBy = "artist"
)

type BlackListBy string

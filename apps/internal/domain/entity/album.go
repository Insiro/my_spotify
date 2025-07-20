package entity

import (
	"github.com/Insiro/my_spotify/internal/domain/entity/shared"
	"time"
)

type Album struct {
	Id                   uint           `gorm:"autoIncrement"`
	AlbumType            string         ``
	Name                 string         `gorm:"size:255"`
	Popularity           uint           `gorm:"default:0"`
	ReleaseDate          time.Time      `gorm:"type:date"`
	ReleaseDatePrecision string         `gorm:"size:20"`
	Uri                  string         `gorm:"size:500"`
	Href                 string         `gorm:"size:500"`
	Meta                 shared.ModMeta `gorm:"embedded"`
	AvailableMarkets     []AlbumMarkets `json:"available_markets"`
	Images               []Image        `gorm:"many2many:album_images"`
	Artists              []Artist       `gorm:"many2many:album_artists;index:idx_album_artists"`
	Tracks               []Track        `gorm:"many2many:album_tracks"`

	// CopyRights   []any   `gorm:"-" json:"copy_rights"`
	// ExternalIds  []any   `gorm:"-" json:"external_ids"`
	// ExternalUrls []any   `gorm:"-" json:"external_urls"`
}

type AlbumMarkets struct {
	AlbumId int      `json:"album_id"`
	Markets []string `gorm:"size:20" json:"markets"`
}

type SpotifyAlbum struct {
	Album
	Artists []Artist `json:"artists"`
	Track   []Track  `json:"track"`
}

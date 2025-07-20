package entity

import "github.com/Insiro/my_spotify/internal/domain/entity/shared"

type Track struct {
	Id          uint
	Album       Album          `gorm:"index:idx_track_album"`
	DiskNumber  uint           `gorm:"default:0"`
	TrackNumber uint           `gorm:"default:0"`
	DurationMs  uint           `gorm:"default:0"`
	Explicit    bool           ``
	Popularity  int            `gorm:"default:0"`
	PreviewUrl  string         `gorm:"size:500"`
	Uri         string         `gorm:"size:500"`
	Href        string         `gorm:"size:500"`
	Type        string         `gorm:"size:50"`
	IsLocal     bool           `gorm:"default:false"`
	Meta        shared.ModMeta `gorm:"embedded"`
	Artists     []Artist       `gorm:"many2many:track_artists;index:idx_track_artists"`

	AvailableMarkets []TrackMarkets
	ExternalIds      any
	ExternalUrls     any
	Name             string `gorm:"size:255" `
}

type TrackMarkets struct {
	Track   Track    `gorm:"primaryKey"`
	Markets []string `gorm:"size:20" `
}

type SpotifyTrack struct {
	Artists []Artist
	Album   SpotifyAlbum
}

type RecentlyPlayedTrack struct {
	PlayedAt string `json:"played_at"`
	track    SpotifyTrack
}

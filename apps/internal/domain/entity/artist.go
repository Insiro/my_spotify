package entity

import (
	"github.com/Insiro/my_spotify/internal/domain/entity/shared"
)

type Artist struct {
	Id         uint           `gorm:"autoIncrement"`
	Name       string         `gorm:"size:255"`
	Popularity int            `gorm:"default:0"`
	Uri        string         `gorm:"size:255"`
	Href       string         `gorm:"size:255"`
	Type       string         `gorm:"size:255"`
	Images     []Image        `gorm:"many2many:artist_images"`
	Genres     []Genre        `gorm:"many2many:artist_genres"`
	Albums     []Album        `gorm:"many2many:artist_albums"`
	Tracks     []Track        `gorm:"many2many:artist_tracks"`
	Meta       shared.ModMeta `gorm:"embedded"`

	//ExternalUrls []any //TODO
	//Followers    any   //TODO
}

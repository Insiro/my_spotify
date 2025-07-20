package entity

import (
	"github.com/Insiro/my_spotify/internal/constant/enum"
	"github.com/Insiro/my_spotify/internal/domain/entity/shared"
	"time"
)

type User struct {
	Id              uint           `gorm:"autoIncrement"`
	UserName        string         `gorm:"size:200"`
	Admin           bool           `gorm:"default:false"`
	SpotifyId       *string        `gorm:"size:200;index"`
	ExpireIn        int            `gorm:"type:int;default:0"`
	AccessToken     *string        `gorm:"size:200;default:null"`
	RefreshToken    *string        `gorm:"size:200;default:null"`
	LastTimestamp   int64          `gorm:"default:0"`
	LastImport      *string        `gorm:"default:null"`
	PublicToken     *string        `gorm:"size:200;default:null;index"`
	FirstListenedAt *time.Time     `gorm:"type:datetime;default:null"`
	Settings        *UserSetting   `json:"settings"`
	Meta            shared.ModMeta `gorm:"embedded"`
	Tracks          []Track        `gorm:"many2many:user_tracks"`
	PlayHistories   []PlayHistory
}

type UserSetting struct {
	User                 User              `gorm:"primaryKey"`
	HistoryLine          bool              `default:"true" gorm:"default:true"`
	PreferredStatsPeriod string            `default:"day" gorm:"default:'day'"`
	NbElements           int               `default:"10" gorm:"default:10" `
	MetricUsed           enum.MetricType   `default:"number" gorm:"default:'number';type:enum('number', 'duration')"`
	DarkMode             enum.DarkModeType `default:"follow" gorm:"default:'follow';type:enum('follow', 'dark', 'light')" `
	Timezone             *string           `gorm:"size(10)"`
	DateFormat           string            `gorm:"size(50)" `
	Meta                 shared.ModMeta    `gorm:"embedded"`
}

type BlackListedArtists struct {
	Id     uint
	User   User
	Artist Artist
	Meta   shared.ModMeta `gorm:"embedded"`
}

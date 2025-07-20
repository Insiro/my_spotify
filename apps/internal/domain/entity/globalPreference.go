package entity

import "github.com/Insiro/my_spotify/internal/domain/entity/shared"

type GlobalPreference struct {
	Id                 uint           `gorm:"autoIncrement"`
	AllowRegistrations bool           `default:"true" gorm:"default:true"`
	AllowAffinity      bool           `default:"true" gorm:"default:true"`
	Meta               shared.ModMeta `gorm:"embedded"`
}

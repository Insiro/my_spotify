package entity

import "github.com/Insiro/my_spotify/internal/domain/entity/shared"

type PrivateData struct {
	Id            uint           `gorm:"autoIncrement"`
	JwtPrivateKey string         `gorm:"primaryKey"`
	Meta          shared.ModMeta `gorm:"embedded"`
}

package shared

import "time"

type ModMeta struct {
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now() on update now()"`
}

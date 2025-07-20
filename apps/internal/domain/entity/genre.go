package entity

type Genre struct {
	Id        int    `gorm:"primary_key"`
	GenreName string `gorm:"size:50;not null;uniqueIndex;"`
}

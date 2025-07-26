package repository

import (
	"github.com/Insiro/my_spotify/internal/domain/entity"
	"github.com/Insiro/my_spotify/pkg"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PrivateData struct {
	db *gorm.DB
}

func (r PrivateData) CreatePrivateData() error {

	token, err := pkg.Crypt.GenerateRandomString(32)

	if err != nil {
		return err
	}

	prvData := entity.PrivateData{JwtPrivateKey: token}
	r.db.Create(&prvData)
	return nil

}

func (r PrivateData) GetPrivateData() (*entity.PrivateData, error) {
	privateData := &entity.PrivateData{}
	result := r.db.Last(&privateData)
	if result.Error != nil {
		return privateData, errors.Wrap(result.Error, "No Private Data Found")

	}

	return privateData, nil
}

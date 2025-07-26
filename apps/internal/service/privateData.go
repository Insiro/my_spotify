package service

import (
	"github.com/Insiro/my_spotify/internal/domain/entity"
	"github.com/Insiro/my_spotify/internal/repository"
)

type PrivateDataService struct {
	privateData repository.PrivateData
}

func NewPrivateDataService(privateData repository.PrivateData) *PrivateDataService {
	return &PrivateDataService{privateData: privateData}
}

func (s *PrivateDataService) CreatePrivateData() error {
	return s.privateData.CreatePrivateData()
}

func (s *PrivateDataService) GetPrivateData() (*entity.PrivateData, error) {
	return s.privateData.GetPrivateData()
}

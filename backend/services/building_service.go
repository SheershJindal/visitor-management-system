package services

import (
	"github.com/sheershjindal/visitor-management-system/models"
	"github.com/sheershjindal/visitor-management-system/repositories"
)

type BuildingService struct {
	repo *repositories.BuildingRepository
}

func NewBuildingService(repo *repositories.BuildingRepository) *BuildingService {
	return &BuildingService{repo: repo}
}

func (s *BuildingService) FetchBuildings() ([]models.Building, error) {
	return s.repo.GetAllBuildings()
}

package repositories

import (
	"github.com/sheershjindal/visitor-management-system/models"
	"gorm.io/gorm"
)

type BuildingRepository struct {
	db *gorm.DB
}

func NewBuildingRepository(db *gorm.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}

func (r *BuildingRepository) GetAllBuildings() ([]models.Building, error) {
	var buildings []models.Building
	if err := r.db.Find(&buildings).Error; err != nil {
		return nil, err
	}
	return buildings, nil
}

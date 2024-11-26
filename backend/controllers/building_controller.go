package controllers

import (
	"net/http"

	"github.com/sheershjindal/visitor-management-system/models"
	"github.com/sheershjindal/visitor-management-system/services"
	"github.com/sheershjindal/visitor-management-system/utils"
)

type BuildingController struct {
	service *services.BuildingService
}

func NewBuildingController(service *services.BuildingService) *BuildingController {
	return &BuildingController{service: service}
}

func (c *BuildingController) GetBuildings(w http.ResponseWriter, r *http.Request) error {
	buildings, err := c.service.FetchBuildings()
	if err != nil {
		return models.NewHTTPError(http.StatusNotFound, "BLD_NOT_FND", err.Error(), nil)
	}

	return utils.SendResponse(w, r, "success", "Fetched Successfully", buildings, nil, http.StatusOK)
}

func (c *BuildingController) CreateBuilding(w http.ResponseWriter, r *http.Request) error {
	return utils.SendResponse(w, r, "success", "Created Successfully", nil, nil, http.StatusCreated)
}

package modules

import (
	"net/http"

	"github.com/sheershjindal/visitor-management-system/controllers"
	"github.com/sheershjindal/visitor-management-system/middlewares"
	"github.com/sheershjindal/visitor-management-system/models"
	"github.com/sheershjindal/visitor-management-system/repositories"
	"github.com/sheershjindal/visitor-management-system/services"
	"gorm.io/gorm"
)

type BuildingModule struct {
	db                 *gorm.DB
	buildingController *controllers.BuildingController
}

func NewBuildingModule(db *gorm.DB) *BuildingModule {
	buildingRepo := repositories.NewBuildingRepository(db)
	buildingService := services.NewBuildingService(buildingRepo)
	buildingController := controllers.NewBuildingController(buildingService)

	return &BuildingModule{
		db:                 db,
		buildingController: buildingController,
	}
}

// GetRouteGroup returns the grouped routes for the Building module
func (m *BuildingModule) GetRouteGroup() models.RouteGroup {
	return models.RouteGroup{
		BasePath: "/buildings",
		Middlewares: []func(http.Handler) http.Handler{
			middlewares.AuthMiddleware,
		},
		Routes: []models.Route{
			{
				Path:    "",
				Method:  http.MethodGet,
				Handler: m.buildingController.GetBuildings,
				// This route doesn't use admin middleware, so it's here instead of a subgroup
			},
		},
		SubGroups: []models.RouteGroup{
			{
				BasePath: "",
				Middlewares: []func(http.Handler) http.Handler{
					middlewares.AuthMiddleware,
				},
				Routes: []models.Route{
					{
						Path:    "",
						Method:  http.MethodPost,
						Handler: m.buildingController.CreateBuilding,
					},
					// {
					// 	Path:    "/{id}",
					// 	Method:  http.MethodPatch,
					// 	Handler: m.buildingController.UpdateBuilding,
					// },
					// {
					// 	Path:    "/{id}",
					// 	Method:  http.MethodDelete,
					// 	Handler: m.buildingController.DeleteBuilding,
					// },
				},
			},
		},
	}
}

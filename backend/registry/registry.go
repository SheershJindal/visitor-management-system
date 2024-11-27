package registry

import (
	"net/http"

	"github.com/sheershjindal/visitor-management-system/config"
	"github.com/sheershjindal/visitor-management-system/middlewares"
	"github.com/sheershjindal/visitor-management-system/models"
	"github.com/sheershjindal/visitor-management-system/modules"
)

// Module interface, defining a GetRoutes function
type Module interface {
	GetRouteGroup() models.RouteGroup
}

// AppRegistry manages modules and their route registration
type AppRegistry struct {
	Modules []Module
}

// NewAppRegistry initializes the registry and all modules internally
func NewAppRegistry(appConfig *config.AppConfig) *AppRegistry {
	return &AppRegistry{
		Modules: []Module{
			modules.NewBuildingModule(appConfig.SQLDB.DB),
			// Add more modules here...
		},
	}
}

// RegisterAllRoutes registers all routes from all modules with the provided mux
func (a *AppRegistry) RegisterAllRoutes(mux *http.ServeMux) {
	// Apply global middlewares
	globalMiddlewares := []func(http.Handler) http.Handler{
		middlewares.LoggingMiddleware,
		middlewares.MetadataMiddleware,
		middlewares.CORSMiddleware,
	}

	// Recursively register routes and groups
	var registerGroup func(basePath string, group models.RouteGroup, middlewares []func(http.Handler) http.Handler)
	registerGroup = func(basePath string, group models.RouteGroup, inheritedMiddlewares []func(http.Handler) http.Handler) {
		// Combine inherited and group-specific middlewares
		groupMiddlewares := append(inheritedMiddlewares, group.Middlewares...)

		// Register individual routes
		for _, route := range group.Routes {
			fullPath := basePath + route.Path
			methodPath := route.Method + " " + fullPath

			handlerWithErrorHandling := middlewares.ErrorRecoveryMiddleware(route.Handler)
			handler := middlewares.ApplyMiddlewares(
				handlerWithErrorHandling,
				append(groupMiddlewares, route.Middlewares...),
			)

			mux.Handle(methodPath, http.HandlerFunc(handler.ServeHTTP))
		}

		// Recursively register subgroups
		for _, subGroup := range group.SubGroups {
			registerGroup(basePath+subGroup.BasePath, subGroup, groupMiddlewares)
		}
	}

	// Register all modules
	for _, module := range a.Modules {
		routeGroup := module.GetRouteGroup()
		registerGroup(routeGroup.BasePath, routeGroup, globalMiddlewares)
	}
}

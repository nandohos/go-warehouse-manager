package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nandohos/go-warehouse-manager/logistics/application/controllers"
)

// TruckRouter struct holds required services for router to function
type TruckRouter struct {
	TruckController controllers.ITruckController
}

// TruckRouterConfig struct holds controllers that will eventually be injected into the router
type TruckRouterConfig struct {
	R               *gin.Engine
	TruckController controllers.ITruckController
}

func NewTruckRouter(c *TruckRouterConfig) {
	// Create router that will have its dependencies injected
	h := &TruckRouter{
		TruckController: c.TruckController,
	}

	// Create group for truck routes
	g := c.R.Group("trucks")

	g.GET("/registration", h.TruckController.Get)
	g.POST("/create", h.TruckController.Create)
}

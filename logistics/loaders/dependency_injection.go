package loaders

import (
	"github.com/gin-gonic/gin"
	"github.com/nandohos/go-warehouse-manager/logistics/application/controllers"
	"github.com/nandohos/go-warehouse-manager/logistics/application/routes"
	"github.com/nandohos/go-warehouse-manager/logistics/domain/services"
	"github.com/nandohos/go-warehouse-manager/logistics/infrastructure/repositories"
)

func Inject(ds *dataSources) (*gin.Engine, error) {

	router := gin.Default()

	// Database Sources
	db := ds.Database

	//Injecting Repositories
	truckRepository := repositories.NewTruckRepository(db)

	//Injecting Services
	truckService := services.NewTruckService(truckRepository)

	//Injecting Controllers
	truckController := controllers.NewTruckController(truckService)

	//Injecting Routes
	routes.NewTruckRouter(&routes.TruckRouterConfig{
		R:               router,
		TruckController: truckController,
	})

	return router, nil

}

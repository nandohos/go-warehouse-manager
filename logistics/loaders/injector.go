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

	db := ds.Database
	truckRepository := repositories.NewTruckRepository(db)
	truckService := services.NewTruckService(truckRepository)
	truckController := controllers.NewTruckController(truckService)
	routes.NewTruckRouter(&routes.TruckRouterConfig{
		R:               router,
		TruckController: truckController,
	})

	return router, nil

}

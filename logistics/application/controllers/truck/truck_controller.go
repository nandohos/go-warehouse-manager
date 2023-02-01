package truck

import (
	"github.com/gin-gonic/gin"
	"github.com/nandohos/go-warehouse-manager/logistics/application/routes"
	"net/http"
)

type TruckRouter struct {
	TruckRouter routes.TruckRouter
}

func (h *TruckRouter) Get(ctx *gin.Context) { //(*dto.Truck, error) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ping",
	})
}

func (h *TruckRouter) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

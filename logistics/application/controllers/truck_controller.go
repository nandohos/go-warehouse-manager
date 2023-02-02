package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nandohos/go-warehouse-manager/logistics/domain/services"
	"net/http"
)

type ITruckController interface {
	Get(ctx *gin.Context)    //(*dto.Truck, error)
	Create(ctx *gin.Context) //error
}

type truckController struct {
	truckService services.ITruckService
}

func NewTruckController(truckService services.ITruckService) ITruckController {
	return &truckController{
		truckService: truckService,
	}
}

func (h *truckController) Get(ctx *gin.Context) { //(*dto.Truck, error) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ping",
	})
}

func (h *truckController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

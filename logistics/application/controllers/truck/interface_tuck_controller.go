package truck

import (
	"github.com/gin-gonic/gin"
)

type ITruckController interface {
	Get(ctx *gin.Context)    //(*dto.Truck, error)
	Create(ctx *gin.Context) //error
}

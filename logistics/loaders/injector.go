package loaders

import "github.com/gin-gonic/gin"

func Inject(ds *dataSources) (*gin.Engine, error) {

	router := gin.Default()

	return router, nil

}


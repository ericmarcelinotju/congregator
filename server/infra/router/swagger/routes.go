package swagger

import (
	"github.com/ericmarcelinotju/congregator/config"
	docs "github.com/ericmarcelinotju/congregator/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRoutesFactory create and returns a factory to create routes for the panelment
func Init(group *gin.RouterGroup) func() {

	config := config.Get()

	docs.SwaggerInfo.Title = config.App.Name
	docs.SwaggerInfo.Description = config.App.Version
	docs.SwaggerInfo.Version = config.App.Version
	docs.SwaggerInfo.Host = config.Server.Rest.Host
	docs.SwaggerInfo.BasePath = config.Server.Rest.Path

	routes := func() {
		group.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return routes
}

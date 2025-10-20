package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/ericmarcelinotju/congregator/config"
	"github.com/ericmarcelinotju/congregator/module/user"
)

// RestServer
func RestServer(cfg *config.Server) {

	gin.DefaultWriter = log.Writer()

	router := gin.Default()

	router.Use()

	// router.NoRoute(func(c *gin.Context) {
	// 	response.ResponseError(c, errors.New("route not found"), http.StatusNotFound)
	// })

	// healthGroup := router.Group("health")
	// healthModule.NewRoutesFactory(healthGroup)()

	userRepo := user.NewRepository(&fiber.Client{})
	userSvc := user.NewService(userRepo)

	// apiGroup := router.Group("api")

	log.Println("Start Listening to : " + cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"

	"github.com/ericmarcelinotju/congregator/config"
	"github.com/gofiber/fiber/v2"

	userModule "github.com/ericmarcelinotju/congregator/module/user"

	"github.com/ericmarcelinotju/congregator/infra/database"
	"github.com/ericmarcelinotju/congregator/infra/router"
)

// @securityDefinitions.apikey Auth
// @in header
// @name Authorization
func main() {
	// get configuration stucts via .env file
	cfg := config.Get()

	// establish DB connection
	db, err := database.Connect(&cfg.Database.Main)
	if err != nil {
		log.Fatalln("[DATABASE] : ", err)
	}

	userRepo := userModule.NewRepository(&fiber.Client{})
	userSvc := userModule.NewService(userRepo)

	router.NewRestServer(
		cfg.Server.Rest,
	)

}

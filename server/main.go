package main

import (
	"github.com/ericmarcelinotju/congregator/config"

	"github.com/ericmarcelinotju/congregator/infra/router"
)

// @securityDefinitions.apikey Auth
// @in header
// @name Authorization
func main() {
	// get configuration stucts via .env file
	cfg := config.Get()

	// establish DB connection
	// db, err := database.Connect(&cfg.Database.Main)
	// if err != nil {
	// 	log.Fatalln("[DATABASE] : ", err)
	// }

	router.RestServer(
		&cfg.Server.Rest,
	)

}

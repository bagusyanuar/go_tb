package main

import (
	"fmt"

	"github.com/bagusyanuar/go_tb/config"
	"github.com/bagusyanuar/go_tb/router"
)

func main() {
	configuration := config.New()
	database := config.NewConnectionDatabase(configuration)
	appPort := configuration.Get("APP_PORT")
	server := router.BuildSheme(database)
	server.Run(fmt.Sprintf(":%s", appPort))
}

package main

import (
	"log"

	"github.com/bagusyanuar/go_tb/config"
)

func main() {
	configuration := config.New()
	database := config.NewConnectionDatabase(configuration)
	database.AutoMigrate(config.User{})
	database.AutoMigrate(config.Member{})
	log.Println("success migrate database")
}

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
	database.AutoMigrate(config.Category{})
	database.AutoMigrate(config.Subject{})
	database.AutoMigrate(config.Grade{})
	database.AutoMigrate(config.MentorLevel{})
	log.Println("success migrate database")
}

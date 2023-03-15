package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/bagusyanuar/go_tb/config"
	"github.com/bagusyanuar/go_tb/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	database.AutoMigrate(config.Province{})
	database.AutoMigrate(config.City{})
	database.AutoMigrate(config.District{})
	database.AutoMigrate(config.Mentor{})
	database.AutoMigrate(config.ProductCourse{})

	database.SetupJoinTable(&config.Subject{}, "Grades", &config.Grade{})
	log.Println("success migrate database")

	if err := database.AutoMigrate(&config.User{}); err == nil && database.Migrator().HasTable(&config.User{}) {
		if err := database.Where("JSON_SEARCH(roles, 'all', 'admin') IS NOT NULL").First(&domain.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			hash, err := bcrypt.GenerateFromPassword([]byte("administrator"), 13)
			if err != nil {
				panic("failed to create seeder")
			}
			tmpPassword := string(hash)

			role, _ := json.Marshal([]string{"admin"})
			admin := domain.User{
				Email:    "administrator@gmail.com",
				Username: "administrator",
				Password: &tmpPassword,
				Roles:    role,
			}
			if err := database.Create(&admin).Error; err != nil {
				panic("failed to create seeder")
			}
		}
		log.Println("success create admin seeder")
	}
}

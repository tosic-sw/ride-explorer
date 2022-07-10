package db

import (
	"RatingService/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=ratingsdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("ratings")
	db.Migrator().AutoMigrate(&models.Rating{})

	// Seeding db

	for _, rating := range Ratings {
		db.Create(&rating)
	}

	return db
}

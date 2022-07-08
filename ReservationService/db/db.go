package db

import (
	"ReservationService/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=reservationsdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("reservations", "cars", "drivers", "passengers", "user_accounts")
	db.Migrator().AutoMigrate(&models.Reservation{})

	// Seeding db

	for _, reservation := range Reservations {
		db.Create(&reservation)
	}

	return db
}

package db

import (
	"UserService/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=usersdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("admins", "cars", "drivers", "passengers", "user_accounts")
	db.Migrator().AutoMigrate(&models.Admin{}, &models.Driver{}, &models.Passenger{}, &models.UserAccount{}, &models.Car{})

	// Seeding db

	for _, admin := range Admins {
		db.Create(&admin)
	}

	for _, passenger := range Passengers {
		db.Create(&passenger)
	}

	for _, driver := range Drivers {
		db.Create(&driver)
	}

	return db
}

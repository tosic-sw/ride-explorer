package db

import (
	"ComplaintService/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=complaintsdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("complaints")
	db.Migrator().AutoMigrate(&models.Complaint{})

	// Seeding db

	for _, compl := range Complaints {
		db.Create(&compl)
	}

	return db
}

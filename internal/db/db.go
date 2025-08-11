package db

import (
	"fmt"
	"log"

	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	connectioninfo := "host=localhost user=root password=Password123 dbname=meu_banco_go port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connectioninfo), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	fmt.Println("Database connection established successfully")

	fmt.Println("Executing database migrations...")

	err = db.AutoMigrate(&models.User{}, &models.Owner{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	fmt.Println("Database migrations executed successfully")
	fmt.Println("Successfully connected to the database")

	return db, nil
}

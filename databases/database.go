package databases

import (
	"StreefySherehes/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string
	Description string
	Date        string
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	TicketPrice float64
}

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	//load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",

		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),




	)
		// Connect to the database
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("❌ Failed to connect to database: ", err)
			return
		}
	
		fmt.Println("✅ Connected to the database successfully")
	
		// Auto-migrate will create the table if it doesn't exist
		DB.AutoMigrate(&Event{})
		DB.AutoMigrate(&models.Event{})
		DB.AutoMigrate(&models.User{})
	}


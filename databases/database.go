package databases

import (
	"fmt"
	"log"
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

	dsn := "host=localhost user=postgres password=N333336K41777757@2019 dbname=postgres port=5432 sslmode=disable TimeZone=Africa/Nairobi"
		// Connect to the database
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("❌ Failed to connect to database: ", err)
			return
		}
	
		fmt.Println("✅ Connected to the database successfully")
	
		// Auto-migrate will create the table if it doesn't exist
		DB.AutoMigrate(&Event{})
	}


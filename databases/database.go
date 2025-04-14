package databases

import (
	"fmt"
	"log"
	"time"

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

	dsn := "host=localhost user=postgres password=N36K4157@2019 dbname=postgres port=5432 sslmode=disable TimeZone=Africa/Nairobi"
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


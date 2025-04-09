package databases


import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"log"
)

var DB *gorm.DB

func ConnectDatabase(){
	var err error
	dsn := "user=yourusername password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to databse: ", err)

	}
	fmt.Println("Connected to the database successfully")
	// Migrate the schema
}
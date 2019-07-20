package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // need to load postgres dialect
)

var db *gorm.DB

func init() {
	// env := godotenv.Load("./../.env")
	// if env == nil {
	// 	log.Fatal("Failed to load environment")
	// } else {
	// 	log.Print(env)
	// }

	// username := os.Getenv("db_user")
	// password := os.Getenv("db_pass")
	// dbName := os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")
	// dbPort := os.Getenv("db_port")
	username := "postgres"
	password := "mysecretpassword"
	dbName := "todo"
	dbPort := "5433"
	dbHost := "localhost"

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort)
	connection, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	db = connection
	db.AutoMigrate(&ToDo{})
}

// GetDB returns a pointer to the DB connection
func GetDB() *gorm.DB {
	return db
}

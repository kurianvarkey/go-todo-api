package db

import (
	"fmt"
	"log"
	"os"

	"github.com/kurianvarkey/go-todo-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_handler *gorm.DB = nil

// Get the connection
func GetConnection() *gorm.DB {
	if db_handler == nil {
		db_handler = Connect()
	}

	return db_handler
}

// Connect connects go to mysql database
func Connect() *gorm.DB {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", db_user, db_pass, db_host, db_name)

	db_handler, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// enable if you want to see the sql logs
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("Kapi API: Failed to connect mysql database " + db_name)
	}

	db_handler.AutoMigrate(&models.Todo{})

	return db_handler
}

// Disconnect is stopping your connection to mysql database
func Disconnect() {
	if db_handler != nil {
		log.Fatalln("Kapi API: Cannot find the database handler")
	}

	db_conn, err := db_handler.DB()
	if err != nil {
		log.Fatalln("Kapi API: Failed to kill connection from database")
	}

	db_conn.Close()
}

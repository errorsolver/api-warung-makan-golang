package models

import (
	// "time"

	// "fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	uri := os.Getenv("DB_URI")
	// fmt.Println(uri)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// sql, err := db.DB()

	// sql.SetMaxOpenConns(5)
	// sql.SetMaxIdleConns(5)
	// sql.SetConnMaxLifetime(time.Minute * 30)

	if err := db.AutoMigrate(&Product{}, &User{}, &TransactionDetail{}, &Transaction{}); err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}

	DB = db
}

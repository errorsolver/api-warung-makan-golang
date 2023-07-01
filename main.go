package main

import (
	"log"

	"golang-wm-api/models"
	"golang-wm-api/routes"

	// "github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail load file .env")
	}

	r := routes.RouterCollection()

	models.ConnectDatabase()

	r.Run(os.Getenv("HOST"))
}

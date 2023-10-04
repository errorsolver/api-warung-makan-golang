package main

import (
	"golang-wm-api/models"
	"golang-wm-api/routes"

	"github.com/joho/godotenv"

	"os"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		panic("Fail load file .env")
	}

	r := routes.RoutesGroup()

	models.ConnectDatabase()

	runErr := r.Run(os.Getenv("HOST"))
	if runErr != nil {
		panic(runErr)
	}
}

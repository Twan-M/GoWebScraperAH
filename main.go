package main

import (
	"GoScraper/handlers"
	"GoScraper/repositories"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_MODE") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if repositories.Connected() {
		log.Println("Starting Database Connection")
	}

	handlers.TitleScraper()
	handlers.AuthPnlUser()
}

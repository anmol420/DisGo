package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientId := os.Getenv("CLIENT_ID")
	log.Println("Client ID:", clientId)
}

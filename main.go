package main

import (
	"her/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Server is running!")
	godotenv.Load()
	api.Server()
}

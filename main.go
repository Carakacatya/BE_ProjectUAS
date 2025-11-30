package main

import (
	"log"

	"projectuasbe/config"
	"projectuasbe/database"
)

func main() {
	// Load .env
	config.LoadConfig()

	// Init DB
	database.InitPostgres()
	database.InitMongo()

	log.Println("🚀 Server running on port", config.AppConfig.AppPort)

	// TODO: initialize Fiber + routes here
}

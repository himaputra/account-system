package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envFile := ".env.local"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Failed to load file %s: %v", envFile, err)
	}

	app, err := InitOrmApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}

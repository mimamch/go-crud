package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/mimamch/go-crud/internal/handlers"
	"github.com/mimamch/go-crud/internal/server"
)

func main() {
	godotenv.Load()
	config := server.Config{
		DBUser: os.Getenv("DB_USERNAME"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_DATABASE"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
	}
	s, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	// go models.InitModels(s.DB)

	handlers.RegisterUserHandler(s)

	s.Run(fmt.Sprintf("%v:%v", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT")))
}

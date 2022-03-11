package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"todos/api"
	"todos/resources"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	db := resources.NewDataBase()
	api.StartServer(db, port)
}

package main

import (
	"fmt"
	server "mtgbc/server"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	port, portExists := os.LookupEnv("PORT")
	if !portExists {
		port = "4000"
	}

	server.StartNode(":" + port)
}

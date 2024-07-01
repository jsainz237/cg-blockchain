package main

import (
	"fmt"
	bc "mtgbc/blockchain"
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
	blockchain := bc.CreateBlockchain()
	port, portExists := os.LookupEnv("PORT")

	if !portExists {
		port = "4000"
	}

	serverContext := server.ServerContext{
		Blockchain: &blockchain,
	}

	server.StartNode(&serverContext, ":"+port)
}

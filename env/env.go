package env

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var defaults = map[string]string{
	"PORT": "4000",
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	// Parse --port command line flag
	port := flag.String("port", os.Getenv("PORT"), "Port to run the server on")
	flag.Parse()

	if *port != "" && *port != os.Getenv("PORT") {
		os.Setenv("PORT", *port)
	}

	for key, value := range defaults {
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}
}

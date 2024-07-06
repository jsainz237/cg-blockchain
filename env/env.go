package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var defaults = map[string]interface{}{
	"PORT": "4000",
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	for key, value := range defaults {
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value.(string))
		}
	}
}

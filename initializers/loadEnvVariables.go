package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	} else {
		fmt.Println("Loaded .env file.")
	}
}

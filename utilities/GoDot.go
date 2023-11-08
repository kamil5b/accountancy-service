package utilities

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariableTest(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		return GoDotEnvVariableTest(key)
	}

	return os.Getenv(key)
}

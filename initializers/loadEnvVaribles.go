package initializers

import (
	"January_1st_Calculator/logging"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		logger := logging.GetLogger()
		logger.Fatal("Error loading .env file")
		log.Fatal("Error loading .env file")
	}
}

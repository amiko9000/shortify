package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func ParseEnvironmentFile()  {
	parseFile, err := strconv.ParseBool(os.Getenv("PARSE_ENVIRONMENT_FILE"))
	if err != nil {
		log.Fatal("Error reading env to bool")
	}

	if parseFile {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func Config() {
	ParseEnvironmentFile()
}
package credentials

import (
	"os"

	"github.com/joho/godotenv"
)

// Database host
var DB_HOST = ""
var DB_PORT = ""

// Database credentials
var DB_USERNAME = ""
var DB_PASSWORD = ""
var DB_DATABASE = ""

func Load() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")

	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DATABASE = os.Getenv("DB_DATABASE")
}

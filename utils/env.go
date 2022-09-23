package utils

import "github.com/joho/godotenv"

func InitENV() {
	err := godotenv.Load()
	if err != nil {
		Logger().Fatal("Error loading .env file")
	}
}

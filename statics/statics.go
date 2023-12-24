package statics

import (
	"os"
)

var (
	SecretKey  string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBTimeZone string
)

func Read() {
	SecretKey = os.Getenv("SECRET_KEY")
	DBHost = os.Getenv("POSTGRES_HOST")
	DBUser = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName = os.Getenv("POSTGRES_DB")
	DBPort = os.Getenv("POSTGRES_PORT")
	DBSSLMode = os.Getenv("SSLMODE")
	DBTimeZone = os.Getenv("TIMEZONE")
}

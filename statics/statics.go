package statics

import (
	"flag"
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
	if IsTesting() {
		SecretKey = "Ajagh12Vajagh$ecr3tK3Y"
		DBHost = "localhost"
		DBUser = "testuser"
		DBPassword = "testingpass"
		DBName = "testdb"
		DBPort = "5432"
		DBSSLMode = "disable"
		DBTimeZone = "UTC"
	} else {
		SecretKey = os.Getenv("SECRET_KEY")
		DBHost = os.Getenv("POSTGRES_HOST")
		DBUser = os.Getenv("POSTGRES_USER")
		DBPassword = os.Getenv("POSTGRES_PASSWORD")
		DBName = os.Getenv("POSTGRES_DB")
		DBPort = os.Getenv("POSTGRES_PORT")
		DBSSLMode = os.Getenv("SSLMODE")
		DBTimeZone = os.Getenv("TIMEZONE")
	}
}

func IsTesting() bool {
	return flag.Lookup("test.v") != nil
}

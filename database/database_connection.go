package database

import (
	"fmt"

	"github.com/OmidRasouli/amuse-park/statics"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Initialize(postgresModels ...interface{}) {
	var err error

	dsn := config()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL database:", err)
	}

	DB.AutoMigrate(postgresModels...)
}

func config() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		statics.DBHost, statics.DBUser, statics.DBPassword, statics.DBName, statics.DBPort, statics.DBSSLMode, statics.DBTimeZone)
	return dsn
}

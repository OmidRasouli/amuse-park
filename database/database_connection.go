package database

import (
	"fmt"

	"github.com/OmidRasouli/amuse-park/models"
	"github.com/OmidRasouli/amuse-park/statics"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseHandler interface {
	CreateAccount(account *models.Account) error
	CreateAuthentication(authentication *models.Authentication) error
}

type RealDatabaseHandler struct {
	DB *gorm.DB
}

func NewRealDatabaseHandler(db *gorm.DB) *RealDatabaseHandler {
	return &RealDatabaseHandler{DB: db}
}

var (
	dbMigrator gorm.Migrator
	dbHandler  DatabaseHandler
)

func Initialize(postgresModels ...interface{}) (*gorm.DB, error) {
	dsn := config()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL database:", err)
	}

	db.AutoMigrate(postgresModels...)
	dbMigrator = db.Migrator()

	dbHandler = NewRealDatabaseHandler(db)
	return db, nil
}

func config() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		statics.DBHost, statics.DBUser, statics.DBPassword, statics.DBName, statics.DBPort, statics.DBSSLMode, statics.DBTimeZone)
	return dsn
}

func Migrator() gorm.Migrator {
	return dbMigrator
}

func DBHandler() DatabaseHandler {
	return dbHandler
}

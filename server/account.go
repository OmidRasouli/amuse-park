package server

import (
	"errors"
	"time"

	"github.com/OmidRasouli/amuse-park/database"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func createAccount(userAccount UserAccount) (*models.Account, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	account := models.Account{
		UserID:   userID,
		Username: userAccount.Username,
		Role:     "player",
		Profile: &models.Profile{
			Level:       1,
			DisplayName: userAccount.Username,
			CreatedDate: time.Now(),
			TimeZone:    *time.UTC,
			State:       "active",
			Email:       userAccount.Email,
		},
	}

	err = database.DBHandler.CreateAccount(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func CreateAuthentication(account *models.Account, password string) (*models.Authentication, error) {
	hashedPassword, err := hashPassword(password)

	if err != nil {
		return nil, err
	}

	authentication := models.Authentication{
		Username: account.Username,
		Password: hashedPassword,
		DeviceID: account.DeviceID,
	}

	err = database.DBHandler.CreateAuthentication(&authentication)
	if err != nil {
		return nil, err
	}

	return &authentication, nil
}

func hashPassword(password string) (string, error) {
	pass := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("something occurred while generating password")
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, pass)
	if err != nil {
		return "", errors.New("password doesn't match with hashed password")
	}

	return string(hashedPassword), nil
}

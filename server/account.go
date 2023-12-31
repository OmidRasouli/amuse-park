package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/OmidRasouli/amuse-park/database"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func createAccount(userAccount *UserAccount) (*models.Account, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	account := models.Account{
		UserID:    userID,
		Username:  userAccount.Username,
		Role:      "player",
		ProfileID: userID,
		Profile: &models.Profile{
			ID:          userID,
			Level:       1,
			DisplayName: userAccount.Username,
			CreatedDate: time.Now(),
			TimeZone:    time.UTC.String(),
			State:       "active",
			Email:       userAccount.Email,
		},
	}

	err = database.DBHandler().CreateAccount(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func updateProfile(profile models.Profile) error {
	dbProfile, err := database.DBHandler().GetProfile(profile.ID)
	if err != nil {
		return fmt.Errorf("something occurred while getting profile: %v", err)
	}

	err = database.DBHandler().UpdateProfile(dbProfile, &profile)
	if err != nil {
		return fmt.Errorf("something occurred while updating profile: %v", err)
	}
	return nil
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

	err = database.DBHandler().CreateAuthentication(&authentication)
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

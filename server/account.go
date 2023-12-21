package server

import (
	"time"

	"github.com/OmidRasouli/amuse-park/database"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/google/uuid"
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
		},
	}

	if err := database.CreateAccount(&account); err != nil {
		return nil, err
	}
	return &account, nil
}

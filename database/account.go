package database

import (
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/google/uuid"
)

func (handler *RealDatabaseHandler) CreateAccount(account *models.Account) error {
	result := handler.DB.Create(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (handler *RealDatabaseHandler) CreateAuthentication(authentication *models.Authentication) error {
	result := handler.DB.Create(authentication)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (handler *RealDatabaseHandler) GetProfile(uuid uuid.UUID) (*models.Profile, error) {
	profile := &models.Profile{}
	result := handler.DB.First(&profile, "id = ?", uuid)
	if result.Error != nil {
		return nil, result.Error
	}
	return profile, nil
}

func (handler *RealDatabaseHandler) UpdateProfile(existingProfile *models.Profile, updateProfile *models.Profile) error {
	result := handler.DB.Model(existingProfile).Updates(updateProfile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

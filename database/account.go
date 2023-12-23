package database

import "github.com/OmidRasouli/amuse-park/models"

func (r *RealDatabaseHandler) CreateAccount(account *models.Account) error {
	result := r.DB.Create(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RealDatabaseHandler) CreateAuthentication(authentication *models.Authentication) error {
	result := r.DB.Create(authentication)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

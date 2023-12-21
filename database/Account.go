package database

import "github.com/OmidRasouli/amuse-park/models"

func CreateAccount(account *models.Account) error {
	result := DB.Create(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

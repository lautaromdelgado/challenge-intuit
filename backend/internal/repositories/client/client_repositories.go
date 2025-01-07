package clients_repositories

import (
	"challenge-intuit/database"
	clients_models "challenge-intuit/internal/models/clients"
)

func GetClientByID(clientid uint) (*clients_models.Client, error) {
	client := new(clients_models.Client)
	db := database.GetDB()
	if err := db.Where("id = ?", clientid).Preload("Domicilio").Preload("NombresApellidos").Find(&client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

package clients_repositories

import (
	"challenge-intuit/database"                               // Importar el paquete database
	clients_models "challenge-intuit/internal/models/clients" // Importar el modelo de clientes
)

// GetClientByID obtiene un cliente mediante el ID
func GetClientByID(clientid uint) (*clients_models.Client, error) {
	client := new(clients_models.Client)
	db := database.GetDB()
	if err := db.Where("id = ?", clientid).Preload("Domicilio").Preload("NombresApellidos").Find(&client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

// GetAllClients obtiene todos los clientes
func GetAllClients() ([]clients_models.Client, error) {
	clients := []clients_models.Client{}
	db := database.GetDB()
	if err := db.Preload("Domicilio").Preload("NombresApellidos").Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

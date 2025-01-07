package clients_services

import (
	clients_models "challenge-intuit/internal/models/clients" // Importar el modelo de clientes
	clients "challenge-intuit/internal/repositories/client"   // Importar el repositorio de clientes
)

// GetClientByID obtiene un cliente mediante el ID
func GetClientByID(clientid uint) (*clients_models.Client, error) {
	client, err := clients.GetClientByID(clientid)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetAllClients obtiene todos los clientes
func GetAllClients() ([]clients_models.Client, error) {
	clients, err := clients.GetAllClients()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

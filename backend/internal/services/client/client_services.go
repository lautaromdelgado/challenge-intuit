package clients_services

import (
	clients_models "challenge-intuit/internal/models/clients"
	clients "challenge-intuit/internal/repositories/client"
)

func GetClientByID(clientid uint) (*clients_models.Client, error) {
	client, err := clients.GetClientByID(clientid)
	if err != nil {
		return nil, err
	}
	return client, nil
}

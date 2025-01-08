package clients_repositories

import (
	"challenge-intuit/database"                                           // Importar el paquete database
	clients_models "challenge-intuit/internal/models/clients"             // Importar el modelo de clientes
	domicilio_models "challenge-intuit/internal/models/domicilio"         // Importar el modelo de domicilio
	nombres_apellido "challenge-intuit/internal/models/nombres_apellidos" // Importar el modelo de nombres y apellidos
)

// OBTENER CLIENTE
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

// CREAR CLIENTE
// CreateClient crea un cliente
func CreateClient(client *clients_models.Client) error {
	db := database.GetDB()
	if err := db.Save(&client).Error; err != nil {
		return err
	}
	return nil
}

// CreateDomicilio crea un domicilio
func CreateDomicilio(domicilio *domicilio_models.Domicilio) error {
	db := database.GetDB()
	if err := db.Save(&domicilio).Error; err != nil {
		return err
	}
	return nil
}

// CreateNombresApellidos crea un nombre y apellido
func CreateNombreApellido(nombreApellido *nombres_apellido.NombresApellidos) error {
	db := database.GetDB()
	if err := db.Save(&nombreApellido).Error; err != nil {
		return err
	}
	return nil
}

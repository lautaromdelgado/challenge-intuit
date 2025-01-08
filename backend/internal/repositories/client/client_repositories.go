package clients_repositories

import (
	"challenge-intuit/database"                                           // Importar el paquete database
	clients_models "challenge-intuit/internal/models/clients"             // Importar el modelo de clientes
	domicilio_models "challenge-intuit/internal/models/domicilio"         // Importar el modelo de domicilio
	nombres_apellido "challenge-intuit/internal/models/nombres_apellidos" // Importar el modelo de nombres y apellidos
	"errors"
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

// UpdateClient actualiza un cliente
func UpdateClient(client *clients_models.Client) error {
	db := database.GetDB()
	if err := db.Save(&client); err != nil {
		return errors.New("Error al actualizar el cliente") // Error personalizado
	}
	return nil
}

// UpdateDomicilio actualiza un domicilio
func UpdateDomicilio(id uint, domicilio *domicilio_models.Domicilio) error {
	db := database.GetDB()
	if err := db.Model(&domicilio).Where("id = ?", id).Updates(&domicilio).Error; err != nil {
		return errors.New("Error al actualizar el domicilio") // Error personalizado
	}
	return nil
}

// UpdateNombreApellido actualiza un nombre y apellido
func UpdateNombreApellido(id uint, nombreApellido *nombres_apellido.NombresApellidos) error {
	db := database.GetDB()
	if err := db.Model(&nombreApellido).Where("id = ?", id).Updates(&nombreApellido).Error; err != nil {
		return errors.New("Error al actualizar el nombre y apellido") // Error personalizado
	}
	return nil
}

// SearchClients busca clientes
func SearchClients(search string) ([]clients_models.Client, error) {
	db := database.GetDB()

	clients := []clients_models.Client{}
	if err := db.Preload("NombresApellidos").Preload("Domicilio").
		Joins("JOIN nombres_apellidos_client ON clientes.id_nombres_apellidos = nombres_apellidos_client.id").
		Where("nombres_apellidos_client.first_name LIKE ? OR nombres_apellidos_client.second_name LIKE ? OR nombres_apellidos_client.first_surname LIKE ? OR nombres_apellidos_client.second_surname LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

package clients_services

import (
	clients_models "challenge-intuit/internal/models/clients"                     // Importar el modelo de clientes
	domicilio_models "challenge-intuit/internal/models/domicilio"                 // Importar el modelo de domicilio
	nombres_apellidos_models "challenge-intuit/internal/models/nombres_apellidos" // Importar el modelo de nombres y apellidos
	clients "challenge-intuit/internal/repositories/client"                       // Importar el repositorio de clientes
	"log"

	"github.com/labstack/echo/v4"
)

// ResponseMessage estructura de respuesta
type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

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

// CreateClient crea un cliente
func CreateClient(c echo.Context) error {
	var req clients_models.Client // Estructura temporal para el request

	// Enlazar los datos del cliente
	if err := c.Bind(&req); err != nil {
		return err
	}

	// Imprimir el contenido de req
	log.Printf("Request recibido: %+v\n", req)                       // Imprime la estructura completa
	log.Printf("Fecha de nacimiento: %v\n", req.Fecha_de_nacimiento) // Imprime la fecha específicamente

	// Crear el domicilio
	newDomicilio, err := CreateDomicilio(&req.Domicilio)
	if err != nil {
		return err
	}

	// Crear el nombre y apellido
	newNombresApellidos, err := CreateNombresApellidos(&req.NombresApellidos)
	if err != nil {
		return err
	}

	// Crear el cliente
	newClient := clients_models.Client{
		ID_Nombres_Apellidos: newNombresApellidos.ID,
		ID_Domicilio:         newDomicilio.ID,
		Fecha_de_nacimiento:  req.Fecha_de_nacimiento,
		Cuit:                 req.Cuit,
		Telefono:             req.Telefono,
		Email:                req.Email,
	}

	// Crear el cliente
	err = clients.CreateClient(&newClient)
	if err != nil {
		return err
	}

	return nil
}

// CreateDomicilio crea un domicilio
func CreateDomicilio(domicilio *domicilio_models.Domicilio) (*domicilio_models.Domicilio, error) {
	// Crear el domicilio
	err := clients.CreateDomicilio(domicilio)
	if err != nil {
		return nil, err
	}

	return domicilio, nil
}

// CreateNombresApellidos crea un nombre y apellido
func CreateNombresApellidos(nombresApellidos *nombres_apellidos_models.NombresApellidos) (*nombres_apellidos_models.NombresApellidos, error) {
	// Crear el nombre y apellido
	err := clients.CreateNombreApellido(nombresApellidos)
	if err != nil {
		return nil, err
	}
	return nombresApellidos, nil
}

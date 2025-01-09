package clients_services

import (
	clients_models "challenge-intuit/internal/models/clients"                     // Importar el modelo de clientes
	domicilio_models "challenge-intuit/internal/models/domicilio"                 // Importar el modelo de domicilio
	nombres_apellidos_models "challenge-intuit/internal/models/nombres_apellidos" // Importar el modelo de nombres y apellidos
	clients "challenge-intuit/internal/repositories/client"                       // Importar el repositorio de clientes
	utils "challenge-intuit/utils"                                                // Importar el paquete de utilidades
	"errors"
	"time"

	"log"
	"strconv"

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
	client, err := clients.GetClientByID(clientid) // Obtener el cliente
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetAllClients obtiene todos los clientes
func GetAllClients() ([]clients_models.Client, error) {
	clients, err := clients.GetAllClients() // Obtener todos los clientes
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

	newDomicilio, err := CreateDomicilio(&req.Domicilio) // Crear el domicilio
	if err != nil {
		return err
	}

	newNombresApellidos, err := CreateNombresApellidos(&req.NombresApellidos) // Crear el nombre y apellido
	if err != nil {
		return err
	}

	newClient := clients_models.Client{ // Estructura del cliente
		ID_Nombres_Apellidos: newNombresApellidos.ID,
		ID_Domicilio:         newDomicilio.ID,
		Fecha_de_nacimiento:  req.Fecha_de_nacimiento,
		Cuit:                 req.Cuit,
		Telefono:             req.Telefono,
		Email:                req.Email,
	}

	err = clients.CreateClient(&newClient) // Crear el cliente
	if err != nil {
		return err
	}

	return nil
}

// CreateDomicilio crea un domicilio
func CreateDomicilio(domicilio *domicilio_models.Domicilio) (*domicilio_models.Domicilio, error) {
	err := clients.CreateDomicilio(domicilio) // Crear el domicilio
	if err != nil {
		return nil, err
	}

	return domicilio, nil
}

// CreateNombresApellidos crea un nombre y apellido
func CreateNombresApellidos(nombresApellidos *nombres_apellidos_models.NombresApellidos) (*nombres_apellidos_models.NombresApellidos, error) {
	err := clients.CreateNombreApellido(nombresApellidos) // Crear el nombre y apellido
	if err != nil {
		return nil, err
	}
	return nombresApellidos, nil
}

// UpdateClient actualiza un cliente
func UpdateClient(c echo.Context) error {
	idParam := c.Param("id")         // Obtener el ID del cliente
	id, err := strconv.Atoi(idParam) // Convertir el ID a entero
	if err != nil {
		return err
	}

	clientDB, err := GetClientByID(uint(id)) // Obtener el cliente de la base de datos
	if err != nil {
		return err
	}

	var clientNew clients_models.Client // Estructura temporal para el request
	if err := c.Bind(&clientNew); err != nil {
		return err
	}

	if !utils.IsEmptyDomicilio(&clientNew.Domicilio) { // Si el domicilio no está vacío
		if clientNew.Domicilio.Calle != "" {
			clientDB.Domicilio.Calle = clientNew.Domicilio.Calle
		}
		if clientNew.Domicilio.Numero != "" {
			clientDB.Domicilio.Numero = clientNew.Domicilio.Numero
		}
		if clientNew.Domicilio.Piso != "" {
			clientDB.Domicilio.Piso = clientNew.Domicilio.Piso
		}
		if clientNew.Domicilio.Departamento != "" {
			clientDB.Domicilio.Departamento = clientNew.Domicilio.Departamento
		}
		if clientNew.Domicilio.Ciudad != "" {
			clientDB.Domicilio.Ciudad = clientNew.Domicilio.Ciudad
		}
		if clientNew.Domicilio.Provincia != "" {
			clientDB.Domicilio.Provincia = clientNew.Domicilio.Provincia
		}
		if clientNew.Domicilio.Codigo_postal != "" {
			clientDB.Domicilio.Codigo_postal = clientNew.Domicilio.Codigo_postal
		}
		if clientNew.Domicilio.Pais != "" {
			clientDB.Domicilio.Pais = clientNew.Domicilio.Pais
		}

		err := clients.UpdateDomicilio(clientDB.Domicilio.ID, &clientDB.Domicilio) // Actualizar el domicilio
		if err != nil {
			return err
		}
	}

	if !utils.IsEmptyNombresApellidos(&clientNew.NombresApellidos) { // Si el nombre y apellido no está vacío
		if clientNew.NombresApellidos.First_name != "" {
			clientDB.NombresApellidos.First_name = clientNew.NombresApellidos.First_name
		}
		if clientNew.NombresApellidos.Second_name != "" {
			clientDB.NombresApellidos.Second_name = clientNew.NombresApellidos.Second_name
		}
		if clientNew.NombresApellidos.First_surname != "" {
			clientDB.NombresApellidos.First_surname = clientNew.NombresApellidos.First_surname
		}
		if clientNew.NombresApellidos.Second_surname != "" {
			clientDB.NombresApellidos.Second_surname = clientNew.NombresApellidos.Second_surname
		}

		err := clients.UpdateNombreApellido(clientDB.NombresApellidos.ID, &clientNew.NombresApellidos) // Actualizar el nombre y apellido
		if err != nil {
			return err
		}
	}

	// Actualizar el cliente
	if clientNew.Fecha_de_nacimiento != nil {
		clientDB.Fecha_de_nacimiento = clientNew.Fecha_de_nacimiento
	}
	if clientNew.Cuit != "" {
		clientDB.Cuit = clientNew.Cuit
	}
	if clientNew.Telefono != "" {
		clientDB.Telefono = clientNew.Telefono
	}
	if clientNew.Email != "" {
		clientDB.Email = clientNew.Email
	}

	err = clients.UpdateClient(clientDB) // Actualizar el cliente
	if err != nil {
		return nil
	}

	return nil
}

// UpdateDomicilio actualiza un domicilio
func UpdateDomicilio(id uint, domicilio *domicilio_models.Domicilio) error {
	err := clients.UpdateDomicilio(id, domicilio) // Actualizar el domicilio
	if err != nil {
		return err
	}
	return nil
}

// UpdateNombresApellidos actualiza un nombre y apellido
func UpdateNombresApellidos(id uint, nombresApellidos *nombres_apellidos_models.NombresApellidos) error {
	err := clients.UpdateNombreApellido(id, nombresApellidos) // Actualizar el nombre y apellido
	if err != nil {
		return err
	}
	return nil
}

// SearchClientByName busca un cliente por nombre
func SearchClient(c echo.Context) ([]clients_models.Client, error) {
	search := c.QueryParam("name") // Obtener el nombre a buscar
	if search == "" {              // Si no se ha ingresado un nombre
		return nil, errors.New("No se ha ingresado un nombre para buscar")
	}
	clients, err := clients.SearchClients(search) // Buscar el cliente
	if err != nil {
		return nil, err
	}
	return clients, nil
}

// DeleteClient eliminado lógico de un cliente
func DeleteClient(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}

	client, err := clients.GetClientByID(uint(id)) // Obtener el cliente
	if err != nil {
		return errors.New("Cliente no encontrado")
	}

	timeNow := time.Now()          // Obtener la fecha y hora actual
	client.Eliminado_el = &timeNow // Asignar la fecha y hora actual al campo "eliminado_el"

	err = clients.DeleteClient(client) // Eliminar el cliente
	if err != nil {
		return err
	}

	return nil
}

// GetClientsDeleted obtiene los clientes eliminados
func GetClientsDeleted() ([]clients_models.Client, error) {
	clients, err := clients.GetClientsDeleted() // Obtener los clientes eliminados
	if err != nil {
		return nil, err
	}
	return clients, nil // Retornar los clientes eliminados
}

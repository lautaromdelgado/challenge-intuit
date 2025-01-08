package clients_handlers

import (
	clients_services "challenge-intuit/internal/services/client"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ResponseMessage estructura de respuesta
type ResponseMessage struct {
	Status       string        `json:"status"`
	Data         any           `json:"data,omitempty"`
	TotalClients *TotalClients `json:"totalclients,omitempty"`
	Message      string        `json:"message,omitempty"`
}

type TotalClients struct {
	Total int `json:"totalclients,omitempty"` // Total de clientes
}

// GetClientByID obtiene un cliente mediante su ID
func GetClientByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	client, err := clients_services.GetClientByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Client not found")
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   client,
	})
}

// GetAllClients obtiene todos los clientes
func GetAllClients(c echo.Context) error {
	clients, err := clients_services.GetAllClients() // Obtener todos los clientes
	if err != nil {
		return c.JSON(http.StatusNotFound, "Clients not found")
	}

	clientsTotal := TotalClients{ // Total de clientes
		Total: len(clients),
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status:       "success",
		Data:         clients,       // Datos de los clientes
		TotalClients: &clientsTotal, // Total de clientes
	})
}

// CreateClient crea un cliente
func CreateClient(c echo.Context) error {
	if err := clients_services.CreateClient(c); err != nil { // Crear un cliente
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  "error",
			Message: "Error creating client: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, ResponseMessage{ // Respuesta
		Status:  "success",
		Message: "Client created successfully",
	})
}

// UpdateClient actualiza un cliente
func UpdateClient(c echo.Context) error {
	err := clients_services.UpdateClient(c) // Actualizar un cliente
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  "error",
			Message: "Error updating client: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Message: "Client updated successfully",
	})
}

// SearchClients busca clientes
func SearchClients(c echo.Context) error {
	clients, err := clients_services.SearchClient(c) // Buscar clientes
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseMessage{
			Status:  "error",
			Message: "Clients not found: " + err.Error(),
		})
	}

	clientsTotal := TotalClients{ // Total de clientes
		Total: len(clients),
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status:       "success",
		Data:         clients,
		TotalClients: &clientsTotal,
	})
}

// DeleteClient elimina un cliente
func DeleteClient(c echo.Context) error {
	err := clients_services.DeleteClient(c) // Eliminar un cliente
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  "error",
			Message: "Error deleting client: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Message: "Client deleted successfully",
	})
}

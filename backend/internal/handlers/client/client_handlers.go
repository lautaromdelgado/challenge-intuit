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
	Total int `json:"totalclients,omitempty"`
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

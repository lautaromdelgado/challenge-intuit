package clients_handlers

import (
	clients_services "challenge-intuit/internal/services/client"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ResponseMessage estructura de respuesta
type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
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

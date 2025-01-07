package clients_routes

import (
	ClientsController "challenge-intuit/internal/handlers/client" // Importar el controlador de clientes

	"github.com/labstack/echo/v4" // Importar el framework Echo
)

func ClientRoutes(e *echo.Echo) {
	e.GET("/clients/:id", ClientsController.GetClientByID) // Obtener cliente mediante el ID
	e.GET("/clients", ClientsController.GetAllClients)     // Obtener todos los clientes
}

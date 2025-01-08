package clients_routes

import (
	ClientsController "challenge-intuit/internal/handlers/client" // Importar el controlador de clientes

	"github.com/labstack/echo/v4" // Importar el framework Echo
)

func ClientRoutes(e *echo.Echo) {
	// GET
	e.GET("/clients/:id", ClientsController.GetClientByID)    // Obtener cliente mediante el ID
	e.GET("/clients", ClientsController.GetAllClients)        // Obtener todos los clientes
	e.GET("/clients/search", ClientsController.SearchClients) // Buscar clientes

	// POST
	e.POST("/clients/create", ClientsController.CreateClient) // Crear un cliente

	// PUT
	e.PUT("/clients/update/:id", ClientsController.UpdateClient) // Actualizar un cliente
	e.PUT("/clients/delete/:id", ClientsController.DeleteClient) // Eliminar un cliente
}

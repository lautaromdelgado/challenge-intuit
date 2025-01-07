package routes

import (
	ClientsRoutes "challenge-intuit/routes/clients" // Importar las rutas de los clientes

	"github.com/labstack/echo/v4" // Importar el framework Echo
)

func InitRoutes(e *echo.Echo) {
	ClientsRoutes.ClientRoutes(e) // Rutas para los clientes
}

package main

import (
	DB "challenge-intuit/database"

	Routes "challenge-intuit/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // Crear una nueva instancia de Echo

	DB.InitDataBase()    // Inicializar la base de datos
	Routes.InitRoutes(e) // Inicializar las rutas

	e.Logger.Fatal(e.Start(":7001")) // Iniciar el servidor en el puerto 7001
}

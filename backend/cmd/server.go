package main

import (
	DB "challenge-intuit/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // Crear una nueva instancia de Echo

	DB.InitDataBase() // Inicializar la base de datos

	e.Logger.Fatal(e.Start(":7001")) // Iniciar el servidor en el puerto 7001
}

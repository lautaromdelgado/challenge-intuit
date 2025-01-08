package utils

import (
	domicilio_models "challenge-intuit/internal/models/domicilio"                 // Importar el modelo de domicilio
	nombres_apellidos_models "challenge-intuit/internal/models/nombres_apellidos" // Importar el modelo de nombres y apellidos
)

// isEmptyDomicilio verifica si el domicilio está vacío
func IsEmptyDomicilio(domicilio *domicilio_models.Domicilio) bool {
	return domicilio.Calle == "" && domicilio.Numero == "" && domicilio.Ciudad == "" &&
		domicilio.Provincia == "" && domicilio.Codigo_postal == "" && domicilio.Pais == ""
}

// isEmptyNombresApellidos verifica si los nombres y apellidos están vacíos
func IsEmptyNombresApellidos(nombresapellidos *nombres_apellidos_models.NombresApellidos) bool {
	return nombresapellidos.First_name == "" && nombresapellidos.First_surname == "" &&
		nombresapellidos.Second_name == "" && nombresapellidos.Second_surname == ""
}

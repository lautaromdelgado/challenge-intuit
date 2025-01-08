package clients_models

import (
	domicilio "challenge-intuit/internal/models/domicilio"                        // Importar el modelo de domicilio
	nombres_apellidos_models "challenge-intuit/internal/models/nombres_apellidos" // Importar el modelo de nombres y apellidos
	"time"
)

// Client es la estructura de un cliente con validación
type Client struct {
	ID                   uint                                      `json:"id" gorm:"primaryKey"`
	ID_Nombres_Apellidos uint                                      `json:"id_nombres_apellidos"`
	NombresApellidos     nombres_apellidos_models.NombresApellidos `json:"nombres_apellidos" gorm:"foreignKey:ID_Nombres_Apellidos;references:ID"`
	Fecha_de_nacimiento  *time.Time                                `json:"fecha_de_nacimiento"`
	Cuit                 string                                    `json:"cuit"`
	ID_Domicilio         uint                                      `json:"id_domicilio"`
	Domicilio            domicilio.Domicilio                       `json:"domicilio" gorm:"foreignKey:ID_Domicilio;references:ID"`
	Telefono             string                                    `json:"telefono"`
	Email                string                                    `json:"email"`
	Creado_el            *time.Time                                `json:"registrado_el,omitempty" gorm:"autoCreateTime"`
	Eliminado_el         *time.Time                                `json:"eliminado_el,omitempty"`
}

// TableName devuelve el nombre de la tabla
func (Client) TableName() string {
	return "clientes"
}

package clients_models

import (
	domicilio "challenge-intuit/internal/models/domicilio"
	nombres_apellidos_models "challenge-intuit/internal/models/nombres_apellidos"
	"time"
)

type Client struct {
	ID                  uint                                       `json:"id" gorm:"primaryKey"`
	Nombres_Apellidos   *nombres_apellidos_models.NombresApellidos `json:"nombres_apellidos" gorm:"foreignKey:ID"`
	Fecha_de_nacimiento *time.Time                                 `json:"fecha_de_nacimiento"`
	Cuit                string                                     `json:"cuit"`
	Domicilio           *domicilio.Domicilio                       `json:"domicilio" gorm:"foreignKey:ID"`
	Telefono            string                                     `json:"telefono"`
	Email               string                                     `json:"email"`
}

func (Client) TableName() string {
	return "clientes"
}

package domicilio_models

// Domicilio es la estructura de un domicilio
type Domicilio struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Calle         string `json:"calle"`
	Numero        string `json:"numero"`
	Piso          string `json:"piso,omitempty"`
	Departamento  string `json:"departamento,omitempty"`
	Ciudad        string `json:"ciudad"`
	Provincia     string `json:"provincia"`
	Codigo_postal string `json:"codigo_postal"`
	Pais          string `json:"pais"`
}

// TableName devuelve el nombre de la tabla
func (Domicilio) TableName() string {
	return "domicilio"
}

package domicilio_models

type Domicilio struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Calle         string `json:"calle"`
	Numero        string `json:"numero"`
	Piso          string `json:"piso"`
	Departamento  string `json:"departamento"`
	Ciudad        string `json:"ciudad"`
	Provincia     string `json:"provincia"`
	Codigo_postal string `json:"codigo_postal"`
	Pais          string `json:"pais"`
}

func (Domicilio) TableName() string {
	return "domicilios"
}

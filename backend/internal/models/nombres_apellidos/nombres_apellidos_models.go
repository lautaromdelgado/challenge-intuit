package nombres_apellidos_models

type NombresApellidos struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	First_name     string `json:"first_name"`
	Second_name    string `json:"second_name,omitempty"`
	First_surname  string `json:"first_surname"`
	Second_surname string `json:"second_surname,omitempty"`
}

func (NombresApellidos) TableName() string {
	return "nombres_apellidos_client"
}

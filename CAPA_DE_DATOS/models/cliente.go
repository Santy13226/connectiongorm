package models

type Cliente struct {
	Cedula             string `gorm:"column:cedula"`
	Nombres            string `gorm:"column:nombres"`
	Apellidos          string `gorm:"column:apellidos"`
	DireccionDomicilio string `gorm:"column:direccion_domicilio"`
	NumeroCelular      string `gorm:"column:numero_celular"`
	CorreoElectronico  string `gorm:"column:correo_electronico"`
	Contrasena         string `gorm:"column:contrasena"`
	Edad               int    `gorm:"column:edad"`
	Sexo               string `gorm:"column:sexo"`
}

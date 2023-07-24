package models

import "time"

type Cliente struct {
	IdCliente          int       `gorm:"column:id_cliente"`
	Cedula             string    `gorm:"column:cedula"`
	Nombres            string    `gorm:"column:nombres"`
	Apellidos          string    `gorm:"column:apellidos"`
	DireccionDomicilio string    `gorm:"column:direccion_domicilio"`
	NumeroCelular      string    `gorm:"column:numero_celular"`
	CorreoElectronico  string    `gorm:"column:correo_electronico"`
	Contrasena         string    `gorm:"column:contrasena"`
	FechaNacimiento    time.Time `gorm:"column:fecha_nacimiento"`
	Sexo               string    `gorm:"column:sexo"`
}

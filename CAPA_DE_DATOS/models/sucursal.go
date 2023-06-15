package models

type Sucursal struct {
	Id             int    `gorm:"column:id_sucursal"`
	Nombre         string `gorm:"column:nombre"`
	Direccion      string `gorm:"column:direccion"`
	Telefono       string `gorm:"column:telefono"`
	HoraDeAtencion string `gorm:"column:hora_de_atencion"`
}

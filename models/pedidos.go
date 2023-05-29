package models

import "time"

type Pedido struct {
	ID         int       `gorm:"column:id_pedido"`
	IDCliente  int       `gorm:"column:id_cliente"`
	IDSucursal int       `gorm:"column:id_sucursal"`
	Fecha      time.Time `gorm:"column:fecha"`
	Producto   string    `gorm:"column:producto"`
	Cantidad   int       `gorm:"column:cantidad"`
	Precio     float64   `gorm:"column:precio"`
}

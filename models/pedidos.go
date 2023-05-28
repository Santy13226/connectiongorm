package models

import "time"

type Pedido struct {
	ID         int
	IDCliente  int
	IDSucursal int
	Fecha      time.Time
	Producto   string
	Cantidad   int
	Precio     float64
}


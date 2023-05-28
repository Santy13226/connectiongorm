package models

type Producto struct {
	ID          int
	ItemCodigo  string
	Nombre      string
	Descripcion string
	Stock       int
	PVP         float64
	IDSucursal  int
}

package models

type Producto struct {
	ID          int `gorm:"column:id_producto"`
	ItemCodigo  string `gorm:"column:item_codigo"`
	Nombre      string `gorm:"column:nombre"`
	Descripcion string `gorm:"column:descripcion"`
	Stock       int	`gorm:"column:stock"`
	PVP         float64 `gorm:"column:p_v_p"`
	IDSucursal  int `gorm:"column:id_sucursal"`
}

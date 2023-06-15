package repositories

import (
	//"fmt"
	"log"

	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/models"
	"gorm.io/gorm"
)

type ProductoRepository struct {
	Conn *gorm.DB
}

func NewProductoRepository(conn *gorm.DB) *ProductoRepository {
	return &ProductoRepository{Conn: conn}
}

// Método de registro de producto
func (r *ProductoRepository) RegistrarProducto(producto *models.Producto) error {
	rs := r.Conn.Create(producto)
	if rs.Error != nil {
		return rs.Error
	}
	log.Println("REGISTRO DE PRODUCTO REALIZADO CORRECTAMENTE...")
	return nil
}

// Método de consulta de productos
func (r *ProductoRepository) ConsultarProductos() ([]models.Producto, error) {
	var productos []models.Producto
	if err := r.Conn.Find(&productos).Error; err != nil {
		return nil, err
	}
	log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	return productos, nil
}

// Método de actualización de productos
func (r *ProductoRepository) ActualizarProducto(nombre string, producto *models.Producto) error {
	result := r.Conn.Where("Nombre=?", nombre).Updates(&producto)
	if result.Error != nil {
		return result.Error
	}
	log.Println("REGISTRO ACTUALIZADO CORRECTAMENTE....")
	return nil
}

// Método de eliminación de productos
func (r *ProductoRepository) EliminarProducto(nombre string) error {
	result := r.Conn.Where("Nombre=?", nombre).Delete(&models.Producto{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		log.Println("NO SE ELIMINÓ", nombre)
	} else {
		log.Println("REGISTRO BORRADO CORRECTAMENTE....")
	}
	return nil
}

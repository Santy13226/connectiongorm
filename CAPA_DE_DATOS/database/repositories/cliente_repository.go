package repositories

import (
	"log"

	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/models"

	"gorm.io/gorm"
)

type ClienteRepository struct {
	Conn *gorm.DB
}

// Constructor de ClienteRepository
func NewClienteRepository(conn *gorm.DB) *ClienteRepository {
	return &ClienteRepository{
		Conn: conn,
	}
}

// Método de registro de clientes
func (repo *ClienteRepository) RegistrarCliente(cliente *models.Cliente) error {
	result := repo.Conn.Create(cliente)
	if result.Error != nil {
		return result.Error
	}
	log.Println("REGISTRO DE CLIENTE REALIZADO CORRECTAMENTE...")
	return nil
}

// Método de consulta general de clientes
func (repo *ClienteRepository) ConsultarClientes() ([]models.Cliente, error) {
	var clientes []models.Cliente
	if err := repo.Conn.Find(&clientes).Error; err != nil {
		return nil, err
	}
	log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	return clientes, nil
}

// Método de consulta de clientes por cédula
func (repo *ClienteRepository) ConsultarClientePorCedula(cedula string) (*models.Cliente, error) {
	var cliente models.Cliente
	result := repo.Conn.Where("cedula=?", cedula).Limit(1).Find(&cliente)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("CONSULTA REALIZADA CORRECTAMENTE...")

	if result.RowsAffected == 0 {
		log.Println("No se encontró el registro", cedula)
		return nil, nil
	}
	return &cliente, nil
}

// Método de actualización de clientes
func (repo *ClienteRepository) ActualizarCliente(cedula string, cliente *models.Cliente) error {
	result := repo.Conn.Where("Cedula=?", cedula).Updates(cliente)
	if result.Error != nil {
		return result.Error
	}
	log.Println("REGISTRO ACTUALIZADO CORRECTAMENTE....")
	return nil
}

// Método de eliminación de clientes
func (repo *ClienteRepository) EliminarCliente(cedula string) error {
	result := repo.Conn.Where("Cedula=?", cedula).Delete(&models.Cliente{})
	if result.Error != nil {
		return result.Error
	}
	log.Println("REGISTRO BORRADO CORRECTAMENTE....")
	if result.RowsAffected == 0 {
		log.Println("NO SE ELIMINÓ", cedula)
	}
	return nil
}

package repositories

import (
	"log"

	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/models"
	"golang.org/x/crypto/bcrypt"

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

func (repo *ClienteRepository) ClienteExistente(cedula, email, celular string) (string, error) {

	var countCedula, countEmail, countCelular int64

	// Verificar si existe algún registro que coincida con la cédula, el email o el número de celular
	repo.Conn.Model(&models.Cliente{}).Where("cedula = ?", cedula).Count(&countCedula)
	repo.Conn.Model(&models.Cliente{}).Where("correo_electronico = ?", email).Count(&countEmail)
	repo.Conn.Model(&models.Cliente{}).Where("numero_celular = ?", celular).Count(&countCelular)

	// Determinar el tipo de error basado en el primer campo que coincide
	if cedula != "" && countCedula > 0 {
		return "cedula", nil
	} else if celular != "" && countCelular > 0 {
		return "celular", nil
	} else if email != "" && countEmail > 0 {
		return "email", nil
	}

	return "", nil
}

func (repo *ClienteRepository) ConsultarClienteLogin(email, password string) bool {
	// Verificar si existe algún registro que coincida con el correo electrónico
	var cliente models.Cliente
	repo.Conn.Where("correo_electronico = ?", email).First(&cliente)

	// Verificar si se encontró el cliente
	if cliente.Cedula == "" { // Ajusta el tipo de dato de acuerdo al campo ID en el modelo Cliente
		return false
	}

	// Comparar las contraseñas
	err := bcrypt.CompareHashAndPassword([]byte(cliente.Contrasena), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (repo *ClienteRepository) ConsultarCliente(email string) models.Cliente {
	var cliente models.Cliente
	repo.Conn.Where("correo_electronico = ?", email).First(&cliente)

	return cliente
}

func (repo *ClienteRepository) ConsultarClienteEmail(email string) bool {
	// Verificar si existe algún registro que coincida con el correo electrónico
	var cliente models.Cliente
	repo.Conn.Where("correo_electronico = ?", email).First(&cliente)
	// Verificar si se encontró el cliente
	if cliente.Cedula == "" { // Ajusta el tipo de dato de acuerdo al campo ID en el modelo Cliente
		return false
	}
	return true
}

func (repo *ClienteRepository) ObtenerTelefono(email string) string {
	var cliente models.Cliente
	repo.Conn.Select("numero_celular").Where("correo_electronico = ?", email).First(&cliente)
	return cliente.NumeroCelular
}

func (repo *ClienteRepository) ChangePass(pass string) bool {

	var cliente models.Cliente
	result := repo.Conn.Where("cedula = ?", cliente.Cedula).First(&cliente)
	if result.Error != nil {
		log.Println("Error al obtener el cliente:", result.Error)
		return false
	}
	cliente.Contrasena = pass
	result = repo.Conn.Save(&cliente)
	if result.Error != nil {
		log.Println("Error al guardar los cambios en la base de datos:", result.Error)
		return false
	}

	log.Println("Contraseña actualizada correctamente")

	return true
}

package repositories

import (
	"log"

	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/models"
	"gorm.io/gorm"
)

type PedidoRepository struct {
	Conn *gorm.DB
}

func NewPedidoRepository(conn *gorm.DB) *PedidoRepository {
	return &PedidoRepository{Conn: conn}
}

// RegistrarPedido registra un pedido.
func (r *PedidoRepository) RegistrarPedido(pedido *models.Pedido) error {
	rs := r.Conn.Create(pedido)
	if rs.Error != nil {
		return rs.Error
	}
	log.Println("REGISTRO DE PEDIDO REALIZADO CORRECTAMENTE...")
	return nil
}

// ConsultarPedido consulta todos los pedidos.
func (r *PedidoRepository) ConsultarPedido() ([]models.Pedido, error) {
	var pedidos []models.Pedido
	if err := r.Conn.Find(&pedidos).Error; err != nil {
		return nil, err
	}
	log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	return pedidos, nil
}

// ConsultarPedidoPorID busca un pedido por su ID.
func (r *PedidoRepository) ConsultarPedidoPorID(ID int) (*models.Pedido, error) {
	pedido := &models.Pedido{}
	if err := r.Conn.First(pedido, ID).Error; err != nil {
		return nil, err
	}
	log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	return pedido, nil
}

// ActualizarPedido actualiza un pedido por su ID.
func (r *PedidoRepository) ActualizarPedido(ID int, pedido *models.Pedido) error {
	result := r.Conn.Model(&models.Pedido{}).Where("id_pedido = ?", ID).Updates(pedido)
	if result.Error != nil {
		return result.Error
	}
	log.Println("REGISTRO ACTUALIZADO CORRECTAMENTE....")
	return nil
}

// EliminarPedido elimina un pedido por su ID.
func (r *PedidoRepository) EliminarPedido(ID int) error {
	result := r.Conn.Delete(&models.Pedido{}, ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		log.Println("NO SE ELIMINÓ EL PEDIDO CON ID", ID)
	} else {
		log.Println("REGISTRO BORRADO CORRECTAMENTE....")
	}
	return nil
}

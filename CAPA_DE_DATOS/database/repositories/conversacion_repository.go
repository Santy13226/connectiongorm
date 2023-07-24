package repositories

import (
	//	"fmt"
	"fmt"
	"log"

	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/models"
	//"github.com/lib/pq"
	"gorm.io/gorm"
)

type ConversacionRepository struct {
	Conn *gorm.DB
}

// Constructor de ConversacionRepository
func NewConversacionRepository(conn *gorm.DB) *ConversacionRepository {
	return &ConversacionRepository{
		Conn: conn,
	}
}

func (repo *ConversacionRepository) RegistrarConversacion(conversacion *models.Conversacion) (int, error) {
	// Insertar la conversación en la base de datos utilizando GORM
	if err := repo.Conn.Create(conversacion).Error; err != nil {
		return 0, err
	}

	// La variable conversacion ahora tiene el ID generado por la base de datos
	idConversacion := conversacion.IDConversacion

	log.Println("REGISTRO DE CONVERSACION REALIZADO CORRECTAMENTE...")
	return idConversacion, nil
}

// Método de consulta general de conversaciones
func (repo *ConversacionRepository) ConsultarConversaciones() ([]models.Conversacion, error) {
	var conversaciones []models.Conversacion
	if err := repo.Conn.Find(&conversaciones).Error; err != nil {
		return nil, err
	}
	//fmt.Println(conversaciones)
	log.Println("CONSULTA DE CONVERSACIONES REALIZADA CORRECTAMENTE...")
	return conversaciones, nil
}

// Método de consulta de conversaciones por ID de conversación
func (repo *ConversacionRepository) ConsultarConversacionPorID(IDConversacion int) (*models.Conversacion, error) {
	var conversacion models.Conversacion
	result := repo.Conn.Where("id_conversacion=?", IDConversacion).Limit(1).Find(&conversacion)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("CONSULTA DE CONVERSACION REALIZADA CORRECTAMENTE...")

	if result.RowsAffected == 0 {
		log.Println("No se encontró el registro con ID de conversación:", IDConversacion)
		return nil, nil
	}
	return &conversacion, nil
}

// Método de actualización de conversaciones

func (repo *ConversacionRepository) ActualizarConversacion(IDConversacion int, conversacion *models.Conversacion) error {
	// Obtiene la conversación actual por su ID
	conversacionActual, err := repo.ConsultarConversacionPorID(IDConversacion)
	if err != nil {
		return err
	}

	// Verifica si la conversación existe
	if conversacionActual == nil {
		return fmt.Errorf("conversación con ID %d no encontrada", IDConversacion)
	}

	// Actualiza los mensajes de la conversación actual con los nuevos mensajes
	conversacionActual.Mensajes = conversacion.Mensajes

	// Realiza la actualización en la base de datos
	result := repo.Conn.Model(conversacionActual).Where("id_conversacion = ?", IDConversacion).Updates(models.Conversacion{
		Mensajes: conversacion.Mensajes,
	})
	if result.Error != nil {
		return result.Error
	}

	log.Println("REGISTRO DE CONVERSACION ACTUALIZADO CORRECTAMENTE....")
	return nil
}

// Método de eliminación de conversaciones
func (repo *ConversacionRepository) EliminarConversacion(IDConversacion int) error {
	result := repo.Conn.Where("id_conversacion=?", IDConversacion).Delete(&models.Conversacion{})
	if result.Error != nil {
		return result.Error
	}
	log.Println("REGISTRO DE CONVERSACION BORRADO CORRECTAMENTE....")
	if result.RowsAffected == 0 {
		log.Println("NO SE ELIMINÓ el registro con ID de conversación:", IDConversacion)
	}
	return nil
}

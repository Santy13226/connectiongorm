package models

import (
	"time"

	"github.com/lib/pq"
)

type Conversacion struct {
	IDConversacion int            `gorm:"column:id_conversacion;primaryKey;autoIncrement"`
	IDCliente      int            `gorm:"column:id_cliente"`
	Fecha          time.Time      `gorm:"column:fecha"`
	Mensajes       pq.StringArray `gorm:"column:mensajes;type:[]" sql:"type:text[]"`
}

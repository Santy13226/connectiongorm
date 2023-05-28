package models

import "time"

type Conversacion struct {
	ID        int
	IDCliente int
	Fecha     time.Time
	Mensajes  string
}


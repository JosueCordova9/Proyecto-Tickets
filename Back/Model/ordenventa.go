package model

import (
	"time"
)

type Ordenventa struct {
	Id_orden    uint
	Fecha_orden time.Time
	Total_orden float32
	// Cliente        Clientes `gorm:"foreignKey:Id_cliente_per"`
	Id_cliente_per int
}

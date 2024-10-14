package model

import (
	"time"

	"gorm.io/gorm"
)

type Ordenventa struct {
	gorm.Model
	Fecha_orden time.Time
	Total_orden float32
	// Cliente        Clientes `gorm:"foreignKey:Id_cliente_per"`
	Id_cliente_per int
}

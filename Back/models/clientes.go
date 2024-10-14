package model

import "gorm.io/gorm"

type Clientes struct {
	gorm.Model
	Nombre_cliente    string
	Apellido_cliente  string
	Cedula_cliente    string
	Telefono_cliente  string
	Direccion_cliente string
	// Ordenes           []Orden `gorm:"foreignKey:Id_cliente_per"`
}

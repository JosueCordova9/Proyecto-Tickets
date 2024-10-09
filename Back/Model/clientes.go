package model

type Clientes struct {
	Id_cliente        uint
	Nombre_cliente    string
	Apellido_cliente  string
	Cedula_cliente    string
	Telefono_cliente  string
	Direccion_cliente string
	// Ordenes           []Orden `gorm:"foreignKey:Id_cliente_per"`
}

package model

import "gorm.io/gorm"

type Detalleorden struct {
	gorm.Model
	Id_orden_per      int
	Id_producto_per   int
	Cantidad_producto int
}

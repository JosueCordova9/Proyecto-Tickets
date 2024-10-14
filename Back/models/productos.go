package model

import "gorm.io/gorm"

type Productos struct {
	gorm.Model
	Nombre_producto string
	Precio_producto float32
}

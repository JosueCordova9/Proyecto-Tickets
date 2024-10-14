package model

import "gorm.io/gorm"

type Usuarios struct {
	gorm.Model
	Nombre_usuario   string
	Usuario_usuario  string
	Password_usuario string
}

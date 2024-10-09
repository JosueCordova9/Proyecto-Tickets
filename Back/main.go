package main

import (
	model "BackTickets/Model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getUsuarios(db *gorm.DB) ([]model.Usuarios, error) {
	var users []model.Usuarios
	result := db.Find(&users)
	return users, result.Error

}

func main() {

	// Cadena de conexión
	connStr := "postgres://postgres:12345@localhost:5432/tickets"

	//Connection DB
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//inicializa un servidor
	r := gin.Default()

	r.GET("/usuarios", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": users,
		})
	})
	r.Run()

}

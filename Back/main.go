package main

import (
	model "BackTickets/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//GET

func getUsuarios(db *gorm.DB) ([]model.Usuarios, error) {
	var usuarios []model.Usuarios
	result := db.Find(&usuarios)
	return usuarios, result.Error
}

func getClientes(db *gorm.DB) ([]model.Clientes, error) {
	var clientes []model.Clientes
	result := db.Find(&clientes)
	return clientes, result.Error

}

func getProductos(db *gorm.DB) ([]model.Productos, error) {
	var productos []model.Productos
	result := db.Find(&productos)
	return productos, result.Error

}

func getOrdenes(db *gorm.DB) ([]model.Ordenventa, error) {
	var ordenes []model.Ordenventa
	result := db.Find(&ordenes)
	return ordenes, result.Error

}

func getDetalles(db *gorm.DB) ([]model.Detalleorden, error) {
	var detalles []model.Detalleorden
	result := db.Find(&detalles)
	return detalles, result.Error
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

	//GET

	r.GET("/usuarios", func(c *gin.Context) {
		usuarios, err := getUsuarios(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"usuarios": usuarios,
		})
	})

	r.GET("/clientes", func(c *gin.Context) {
		clientes, err := getClientes(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"clientes": clientes,
		})
	})

	r.GET("/productos", func(c *gin.Context) {
		productos, err := getProductos(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"productos": productos,
		})
	})

	r.GET("/ordenes", func(c *gin.Context) {
		ordenes, err := getOrdenes(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"ordenes": ordenes,
		})
	})

	r.GET("/detalles", func(c *gin.Context) {
		detalles, err := getDetalles(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"detalles": detalles,
		})
	})

	//POST

	r.POST("/usuarios", func(c *gin.Context) {
		var usuario model.Usuarios
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		db.Create(&usuario)

		c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado"})
	})

	db.AutoMigrate(&model.Productos{})

	r.POST("/productos", func(c *gin.Context) {
		var producto model.Productos
		if err := c.ShouldBindJSON(&producto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		db.Create(&producto)

		c.JSON(http.StatusCreated, gin.H{"message": "Producto creado"})
	})
	r.Run()

}

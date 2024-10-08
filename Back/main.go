package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//inicializa un servidor
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Servidor Web con Gin",
		})
	})
	r.Run()

}

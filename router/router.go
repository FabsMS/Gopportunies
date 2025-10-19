package router

import "github.com/gin-gonic/gin"

func Initializer() {
	// Inicia o Router, utilizando as configurações Default do Gin
	router := gin.Default()

	// Define a Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Comando para rodar a API
	router.Run(":8080")
}

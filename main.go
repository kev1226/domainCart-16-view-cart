package main

import (
	"get-cart/config"
	"get-cart/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa conexión a Redis
	config.InitRedis()

	// Crea instancia de Gin
	r := gin.Default()

	// Registra las rutas del carrito
	routes.RegisterCartRoutes(r)

	// Inicia el servidor
	if err := r.Run(":3036"); err != nil {
		log.Fatal("❌ Error starting server: ", err)
	}
}

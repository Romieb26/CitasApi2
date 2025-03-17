package main

import (
	"log"

	"notificaciones/src/core"
	"notificaciones/src/notificaciones/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	core.InitDB()

	// Crear un router con Gin
	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Ajusta el puerto según sea necesario
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Inicializar dependencias
	citaRouter := infrastructure.NewCitaRouter(router)
	citaRouter.Run() // Agregar rutas

	// Iniciar el servidor
	log.Println("Servidor corriendo en http://localhost:8001")
	if err := router.Run(":8001"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}

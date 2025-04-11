// main.go
package main

import (
	"log"

	"notificaciones/src/core"
	"notificaciones/src/notificaciones/application"
	"notificaciones/src/notificaciones/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar conexiones
	core.InitDB()
	core.InitRabbitMQ()
	defer core.CloseChannels()

	// Configurar WebSocket Hub
	hub := infrastructure.NewHub()
	go hub.Run()

	// Crear router Gin
	router := gin.Default()

	// Configurar CORS para WebSocket
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Registrar rutas REST
	citaRouter := infrastructure.NewCitaRouter(router)
	citaRouter.Run()

	// Registrar ruta WebSocket
	router.GET("/ws", func(c *gin.Context) {
		hub.HandleWebSocket(c)
	})

	// Iniciar consumidor RabbitMQ con WebSocket zd
	go func() {
		err := core.ConsumeMessages("citas_creadas", func(body []byte) {
			application.ProcessCitaMessage(body, hub) // Pasamos el hub aquí
		})
		if err != nil {
			log.Fatalf("Error al consumir mensajes: %v", err)
		}
	}()

	// Iniciar servidor
	log.Println("Servidor iniciado en :8001")
	if err := router.Run(":8001"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}

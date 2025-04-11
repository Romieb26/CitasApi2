package infrastructure

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]bool),
	}
}

// SendMessage envía un mensaje a todos los clientes conectados
func (h *Hub) SendMessage(message []byte) {
	h.broadcast <- message
}

// Run inicia el hub para manejar conexiones WebSocket
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("Nuevo cliente WebSocket conectado. Total: %d", len(h.clients))

			// Envía un mensaje de bienvenida
			welcomeMsg := []byte("Conexión WebSocket establecida con el servidor")
			if err := client.WriteMessage(websocket.TextMessage, welcomeMsg); err != nil {
				log.Printf("Error enviando mensaje de bienvenida: %v", err)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
				log.Printf("Cliente WebSocket desconectado. Total: %d", len(h.clients))
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Printf("Error enviando mensaje a cliente: %v", err)
					client.Close()
					delete(h.clients, client)
				}
			}
		}
	}
}

// HandleWebSocket maneja las conexiones WebSocket entrantes
func (h *Hub) HandleWebSocket(c *gin.Context) {
	// Configuración dinámica de CheckOrigin
	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")

		// Permitir todos los orígenes en desarrollo
		if os.Getenv("APP_ENV") == "development" {
			return true
		}

		// Lista de dominios permitidos en producción
		allowedDomains := []string{
			"34.237.191.108", // Tu instancia AWS
			"tudominio.com",  // Reemplaza con tu dominio real
			"localhost",      // Para desarrollo local
		}

		// Verificar si el origen está en la lista de permitidos
		for _, domain := range allowedDomains {
			if strings.Contains(origin, domain) {
				return true
			}
		}

		log.Printf("Intento de conexión desde origen no permitido: %s", origin)
		return false
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error al actualizar a WebSocket: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "No se pudo establecer conexión WebSocket",
			"details": err.Error(),
		})
		return
	}

	// Registrar la nueva conexión
	h.register <- conn

	// Configurar función de limpieza cuando se cierre la conexión
	defer func() {
		h.unregister <- conn
		conn.Close()
	}()

	// Mantener la conexión activa
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error en WebSocket: %v", err)
			}
			break
		}
	}
}

// GetWebSocketURL devuelve la URL correcta del WebSocket
func GetWebSocketURL() string {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		return "ws://34.237.191.108/ws" // Cambia a wss:// si usas SSL
	}
	return "ws://localhost:8001/ws" // Para desarrollo local
}

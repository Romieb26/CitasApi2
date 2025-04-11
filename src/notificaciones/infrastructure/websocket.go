package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Permite todas las conexiones (para desarrollo y pruebas)
		// En producción deberías restringir esto a tus dominios
		return true
	},
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

func (h *Hub) SendMessage(message []byte) {
	h.broadcast <- message
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("Nuevo cliente conectado. Total: %d", len(h.clients))

			welcomeMsg := []byte("Conexión WebSocket establecida")
			if err := client.WriteMessage(websocket.TextMessage, welcomeMsg); err != nil {
				log.Printf("Error enviando mensaje de bienvenida: %v", err)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
				log.Printf("Cliente desconectado. Total: %d", len(h.clients))
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Printf("Error enviando mensaje: %v", err)
					client.Close()
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error al establecer WebSocket: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "No se pudo establecer conexión WebSocket",
			"details": err.Error(),
		})
		return
	}

	h.register <- conn

	defer func() {
		h.unregister <- conn
		conn.Close()
	}()

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

func GetWebSocketURL() string {
	return "ws://34.237.191.108:8001/ws" // Usa esta URL en todos los entornos
}

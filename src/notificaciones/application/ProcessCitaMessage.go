// ProcessCitaMessage.go
package application

import (
	"encoding/json"
	"log"

	"notificaciones/src/notificaciones/domain/entities"
)

// Definición de interfaz para evitar dependencia circular
type MessageBroadcaster interface {
	SendMessage(message []byte)
}

// ProcessCitaMessage procesa los mensajes y envía por WebSocket
func ProcessCitaMessage(message []byte, broadcaster MessageBroadcaster) {
	// Deserializar el mensaje
	var cita entities.Cita
	err := json.Unmarshal(message, &cita)
	if err != nil {
		log.Printf("Error al deserializar el mensaje: %v", err)
		return
	}

	// Convertir a JSON para WebSocket
	jsonData, err := json.Marshal(cita)
	if err != nil {
		log.Printf("Error serializando cita: %v", err)
		return
	}

	// Enviar a todos los clientes WebSocket
	broadcaster.SendMessage(jsonData)
	log.Printf("Cita recibida y enviada por WebSocket: %+v", cita)
}

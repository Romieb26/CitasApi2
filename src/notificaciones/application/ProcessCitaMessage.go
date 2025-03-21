// ProcessCitaMessage.go
package application

import (
	"encoding/json"
	"log"

	"notificaciones/src/notificaciones/domain/entities"
)

// ProcessCitaMessage procesa los mensajes recibidos de la cola "citas_creadas"
func ProcessCitaMessage(message []byte) {
	// Deserializar el mensaje en una estructura Cita
	var cita entities.Cita
	err := json.Unmarshal(message, &cita)
	if err != nil {
		log.Printf("Error al deserializar el mensaje: %v", err)
		return
	}

	// Loggear la cita recibida
	log.Printf("Cita recibida: %+v", cita)
}

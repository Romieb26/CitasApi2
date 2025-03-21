// rabbitmq.go
package core

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection

// InitRabbitMQ inicializa la conexión a RabbitMQ
func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://romina:romina264@3.230.241.180:5672/")
	if err != nil {
		log.Fatalf("No se pudo conectar a RabbitMQ: %v", err)
	}
	RabbitMQConn = conn
	log.Println("Conectado a RabbitMQ")
}

// ConsumeMessages consume mensajes de la cola "citas_creadas"
func ConsumeMessages(queueName string, handler func([]byte)) error {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declarar la cola (por si no existe)
	_, err = ch.QueueDeclare(
		queueName, // nombre de la cola
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Configurar el canal para consumir mensajes
	msgs, err := ch.Consume(
		queueName, // nombre de la cola
		"",        // consumer
		true,      // auto-ack (reconocimiento automático)
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Procesar los mensajes recibidos
	go func() {
		for msg := range msgs {
			log.Printf("Mensaje recibido: %s", msg.Body)
			handler(msg.Body) // Llamar al manejador de mensajes
		}
	}()

	log.Printf("Escuchando mensajes en la cola %s...", queueName)
	return nil
}

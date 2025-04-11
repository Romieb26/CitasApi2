// rabbitmq.go
package core

import (
	"log"
	"sync"

	"github.com/streadway/amqp"
)

var (
	RabbitMQConn *amqp.Connection
	channelMutex sync.Mutex
	channels     []*amqp.Channel // Para manejar múltiples canales si es necesario
)

// InitRabbitMQ inicializa la conexión principal a RabbitMQ
func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://romina:romina264@3.230.241.180:5672/")
	if err != nil {
		log.Fatalf("No se pudo conectar a RabbitMQ: %v", err)
	}
	RabbitMQConn = conn
	log.Println("Conectado a RabbitMQ")
}

// ConsumeMessages consume mensajes de forma persistente
func ConsumeMessages(queueName string, handler func([]byte)) error {
	channelMutex.Lock()
	defer channelMutex.Unlock()

	ch, err := RabbitMQConn.Channel()
	if err != nil {
		return err
	}

	// Guardar el canal para cerrarlo luego
	channels = append(channels, ch)

	// Configurar QoS para evitar sobrecarga
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return err
	}

	// Declarar la cola con características específicas
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queueName,
		"",    // consumerTag (auto-generado)
		false, // autoAck (MANUAL ahora)
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,
	)
	if err != nil {
		return err
	}

	// Goroutine permanente para escuchar mensajes
	go func() {
		for msg := range msgs {
			log.Printf("Mensaje recibido: %s", msg.Body)
			handler(msg.Body)

			// Ack manual SOLO si el procesamiento fue exitoso
			if err := msg.Ack(false); err != nil {
				log.Printf("Error confirmando mensaje: %v", err)
			}
		}
		log.Println("Canal cerrado, reconectando...")
	}()

	log.Printf("Escuchando mensajes en la cola %s...", queueName)
	return nil
}

// CloseChannels cierra todas las conexiones al apagar
func CloseChannels() {
	channelMutex.Lock()
	defer channelMutex.Unlock()

	for _, ch := range channels {
		if err := ch.Close(); err != nil {
			log.Printf("Error cerrando canal: %v", err)
		}
	}

	if RabbitMQConn != nil {
		RabbitMQConn.Close()
	}
}

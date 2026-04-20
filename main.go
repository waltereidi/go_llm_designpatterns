package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 🔌 Conexão com RabbitMQ
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Erro ao conectar no RabbitMQ")
	defer conn.Close()

	// 📡 Canal
	ch, err := conn.Channel()
	failOnError(err, "Erro ao abrir canal")
	defer ch.Close()

	// // 📬 Declarar fila
	// q, err := ch.QueueDeclare(
	// 	"minha_fila", // nome
	// 	true,         // durável
	// 	false,        // auto-delete
	// 	false,        // exclusiva
	// 	false,        // no-wait
	// 	nil,          // args
	// )
	// failOnError(err, "Erro ao declarar fila")

	// 👂 Consumir mensagens
	msgs, err := ch.Consume(
		"minha_fila",
		"",    // consumer
		false, // auto-ack (false = manual)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Erro ao registrar consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("📥 Mensagem recebida: %s", d.Body)

			// ⚙️ Processamento da mensagem
			err := processMessage(d.Body)
			if err != nil {
				log.Printf("❌ Erro ao processar: %s", err)
				d.Nack(false, true) // requeue
				continue
			}

			// ✅ Confirma processamento
			d.Ack(false)
		}
	}()

	log.Println("🚀 Aguardando mensagens...")
	<-forever
}

func processMessage(body []byte) error {
	// Simula lógica de negócio
	log.Printf("🔧 Processando: %s", string(body))
	return nil
}

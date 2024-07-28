package main

import (
	"rabbitmq/impl"
	"rabbitmq/shared"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {

	// cria conexão com o broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao broker")
	defer conn.Close()

	// cria um canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o broker")
	defer ch.Close()

	// declara a fila
	q, err := ch.QueueDeclare(
		shared.RequestQueue,
		false,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no broker")

	// prepara o recebimento de mensagens do cliente
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Falha ao registrar o consumidor no broker")

	fmt.Println("Crivo pronto...")
	for d := range msgs {
		// recebe request
		msg := shared.Request{}
		err := json.Unmarshal(d.Body, &msg)
		shared.ChecaErro(err, "Falha ao desserializar a mensagem")

		// processa request
		r := impl.CrivoDeEratostenes{}.InvocaCrivoDeEratostenes(msg)

		// prepara resposta
		replyMsg := shared.Reply{Result: []interface{}{r}}
		replyMsgBytes, err := json.Marshal(replyMsg)
		shared.ChecaErro(err, "Falha ao serializar mensagem")

		// publica resposta
		err = ch.Publish(
			"",
			d.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: d.CorrelationId, // usa correlation id do request
				Body:          replyMsgBytes,
			},
		)
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o broker")
	}
}
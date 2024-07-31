package main

import (
	"os"
	"rabbitmq/shared"
	"strconv"

	//crivo "rabbitmq/impl"

	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)


func main() {

	// gera nova seed
	rand.Seed(time.Now().UTC().UnixNano())

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao servidor de mensageria")
	defer conn.Close()

	file, err := os.Create("../resultadosRabbitMQ/output_10_100.txt")
	if err != nil {
		fmt.Println("Erro na criação do arquivo:", err)
		return
	}
	defer file.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	// declara a fila para as respostas
	replyQueue, err := ch.QueueDeclare(
		shared.ResponseQueue,
		false,
		false,
		true,
		false,
		nil,
	)

	// cria servidor da fila de response
	msgs, err := ch.Consume(
		replyQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Falha ao registrar o servidor no broker")

	for j := 0; j < shared.StatisticSample; j++ {
		for i := 0; i < 100; i++ {
			t1 := time.Now()

			// prepara mensagem
			msgRequest := shared.Request{Number: 10}
			msgRequestBytes, err := json.Marshal(msgRequest)
			shared.ChecaErro(err, "Falha ao serializar a mensagem")

			correlationID := shared.RandomString(32)

			err = ch.Publish(
				"",
				shared.RequestQueue,
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: correlationID,
					ReplyTo:       replyQueue.Name,
					Body:          msgRequestBytes,
				},
			)
			shared.ChecaErro(err, "Falha ao enviar a mensagem")
			t2 := time.Now().Sub(t1).Nanoseconds()

			// recebe mensagem do servidor de mensageria
			m := <-msgs

			// deserializa e imprime mensagem na tela
			var msgResponse shared.Reply
			err = json.Unmarshal(m.Body, &msgResponse)
			shared.ChecaErro(err, "Erro na deserialização da resposta")

			fmt.Printf("%v\n", msgResponse.Result[0])

			// Converte int64 pra string e salva no arquivo
			_, err = file.WriteString(strconv.FormatInt(t2, 10) + "\n")
			//_, err := fmt.Fprintf(file, "%d\n", strconv.FormatInt(t2, 10))
			if err != nil {
				fmt.Println("Erro na escrita do arquivo:", err)
						}
		}
	}
}

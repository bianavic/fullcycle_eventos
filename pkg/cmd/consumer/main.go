package main

import (
	"fmt"
	"github.com/bianavic/fullcycle_eventos/utils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// abre canal
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// cria canal msg:
	//toda msg enviada para o OUT é recebida aqui nesse canal MSGS
	msgs := make(chan amqp.Delivery)

	// THREAD lendo rabbitmq: roda o consumidor passando o channel do rabbitmq e o out (as mensagens)
	go rabbitmq.Consume(ch, msgs, "orders")

	// toda msg enviada para channel sera pega no for
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		// indicado dar ack false (linha 28) de forma manual para caso a msg seja perdida na linha 26
		msg.Ack(false) // dando um ack indicando que a msg foi lida (processada), é falso pq nao ira colocar na fila novamente
	}
}

package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// criar conexao e canal
func OpenChannel() (*amqp.Channel, error) {

	// criar conexao - dial para a conexao
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	/*
		CANAL rabbitmq
		criar channel - abre canal com rabbitmq (canal do rabbitmq)
	*/
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

/*
CANAL GO
consumir msg que esta na fila
os dados irao sair (msgs com nome out)
*/
func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-consume",
		// auto ack true - qdo recebe a msg da uma baixa - diz q ja foi lida e pode remover da fila - qdo PODE PERDER a msg
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// ler msgs que chegam do consumidor e jogar para canal
	for msg := range msgs {
		out <- msg
	}
	return nil
}

// enviar, publicar mesg
func Publish(ch *amqp.Channel, body string, exchName string) error {
	err := ch.Publish(
		exchName, // ligar a uma fila
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

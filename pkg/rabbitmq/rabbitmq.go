package rabbitmq

import "github.com/rabbitmq/amqp091-go"

// criar conexao e canal
func OpenChannel() (*amqp091.Channel, error) {

	// criar conexao - dial para a conexao
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
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
func Consume(ch *amqp091.Channel, out chan<- amqp091.Delivery) error {
	msgs, err := ch.Consume(
		"minhafila",
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

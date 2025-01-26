package main

import (
	"github.com/bianavic/fullcycle_eventos/utils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// publica para a amq.direct chamada Hello World
	// o consumidor estivera escutando a fila orders (bind)
	rabbitmq.Publish(ch, "Hello World", "amq.direct")
}

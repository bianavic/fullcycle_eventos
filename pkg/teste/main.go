package main

import "fmt"

func main() {
	evento := []string{"teste1", "teste2", "teste3", "teste4"}

	// slice e eventos
	//evento = evento[:2] // [teste1 teste2] mostra ate o segundo registro
	//evento = evento[:0] // [] nao encontra nenhum registro, mesmo considerando o primeiro registro
	//evento = evento[2:] // [teste3 teste4] mostra o restante dos registros, ele INCLUSO
	//evento = evento[1:] // [teste2 teste3 teste4] mostra o restante dos registros, ele EXCLLUIDO - pula o primeiro registro

	/*
		remover primeiro item encontrado no slice e grudar todos os eventos encontrados depois desse primeiro item

		evento[:0] nao encontra nenhum registro
		evento[1:] EXCLLUI o primeiro registro, mostra o restante
	*/
	evento = append(evento[:0], evento[1:]...) // [teste2 teste3 teste4]

	fmt.Println(evento)
}

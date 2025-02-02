## fullcycle_eventos

---

### Biblioteca gerenciadora de eventos
Abaixo estão os elementos taticos do contexto de eventos:

1- Evento: a interface `EventInterface` possui a visao geral do que o sistema faz
  - `GetName` pega dados com o nome do evento
  - `GetDateTime` sabe o momento que evento foi disparado
  - `GetPayload` interface vazia abrangendo diversos payloads em diversos formatos (possui os dados que o evento contem)

2- Operacoes executadas quando o evento é chamado: 
  - o metodo `Handle` é o executor das operacoes geradas pelos eventos e esta contido na interface `EventHandler`

3- Gerenciador: `EventDispatcherInterface` recebe, dispara e executa, remove, verifica e limpa

    - `Register` metodo que registra o evento (recebe o evento e executa)
    - `Dispatch` dispara o evento para que sejam executados: percorre e executa handler por handler
    - `Remove` remove evento
    - `Has` verifica se o evento X possui aquela operacao Y
    - `Clear` limpa o dispacher matando todos os eventos registrados

obs: EventDispatcher - é um `map` porque um evento pode ter diversos handlers registrados

---
### Go Routine
Tornar o método `Dispatch` async: disparar eventos ao mesmo tempo, rodando em paralelo, em threads separadas
- adicionar Go Routine (rodando numa thread separada)
- criar WaitGroup
- criar 1 thread para cada handle
- enqto executa todas as threads, adiciona Wait para esperar a execucao de todos os handlers

---
### RabbitMQ com Docker

acessa pasta
```shell
cd pkg
```

baixar imagem / subir rabbitmq
```shell
docker-compose up -d
```

acessa admin browser
```
http://localhost:15672/
```

rodar consumer
```shell
cd pkg/cmd/consumer
go run main.go
```

rabbitmq configurado
a- consumer:
![consumer.png](assets%3Aimages/consumer.png)

b- fila:
![fila.png](assets%3Aimages/fila.png)

c- publicacao mensagem:
![msg.png](assets%3Aimages/msg.png)

d- msg output:
![msg_output.png](assets%3Aimages/msg_output.png)

e- configura exhange:
![exchange_direct.png](assets%3Aimages/exchange_direct.png)

enviando `hello world` para amq.direct roteando para o consumidor

a- roda consumer
```shell
cd pkg/cmd/consumer
go run main.go
```

b- roda producer
```shell
cd pkg/cmd/producer
go run main.go
```
c- output consumer
![msg_output2.png](assets%3Aimages/msg_output2.png)
### fullcycle_eventos
## Biblioteca gerenciadora de eventos


Elementos taticos do contexto de eventos:

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
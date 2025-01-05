### fullcycle_eventos
## Biblioteca gerenciadora de eventos


Elementos taticos do contexto de eventos:

1- Evento: a interface `EventInterface` que possui a visao geral do que o sistema faz
  - pega dados com o nome do evento
  - sabe o momento que evento foi disparado
  - interface vazia abrangendo diversos payloads em diversos formatos (dados que contem o evento)

2- Operacoes executadas quando o evento é chamado: 
  - a interface `EventHandler` contem o metodo executor da operacao gerada pelo evento

3- Gerenciador: recebe, dispara, executa
    - metodo que registra o evento (recebe o evento e executa)
    - dispara o evento para que sejam executados
    - remove evento
    - verifica se o evento X possui aquela operacao Y
    - limpa o dispacher matando todos os eventos registrados

EventDispatcher: 
1 evento pode ter diversos handlers registrados
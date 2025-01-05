package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	// verifica se existe esse nome de evento dentro do map
	if _, ok := ed.handlers[eventName]; ok {
		// se existir (ok == true) percorro todos os handlers desse evento
		for _, h := range ed.handlers[eventName] {
			// verifica se evento (aquele nome) é igual ao handler que percorro, retorna erro
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	// caso evento nunca tenha sido registrado, adiciona no map
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

package events

import (
	"errors"
	"sync"
)

var (
	ErrHandlerAlreadyRegistered = errors.New("handler already registered")
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

/*
SYNC FUNCTION
*/
//func (ev *EventDispatcher) Dispatch(event EventInterface) error {
//	// verifica se o evento tem um handler registrado com esse nome de evento
//	if handlers, ok := ev.handlers[event.GetName()]; ok {
//		// verifica cada um dos handlers
//		for _, handler := range handlers {
//			// executa o metodo Handle passando o evento que foi chamado
//			handler.Handle(event)
//		}
//		return nil
//	}
//	return errors.New("no handlers registered for event: " + event.GetName())
//}

func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	// verifica tem um handler registrado com esse nome de evento
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		// verifica cada um dos handlers
		for _, handler := range handlers {
			// executa o metodo Handle passando o evento que foi chamado
			wg.Add(1) // precisa adicionar numero 1 para que cada vez que o handler terminar a execucao, adiciona wg.Done
			go handler.Handle(event, wg)
		}
		wg.Wait()
		return nil
	}
	return errors.New("no handlers registered for event: " + event.GetName())
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

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	// verificar se tem registro do evento
	if _, ok := ed.handlers[eventName]; ok {
		// percorrer todos os handlers que estao nos registros desse evento
		for i, h := range ed.handlers[eventName] {
			// qdo encontrar o handler
			if h == handler {
				// fara o append ZERO - o primeiro item que encontra - com o segundo e itens restantes
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	// verifica se o nome do evento esta registrado
	if _, ok := ed.handlers[eventName]; ok {
		// verifica se o evento registrado pertence ao handler
		for _, h := range ed.handlers[eventName] {
			// comparar o handler passado com o que esta registrado
			if h == handler {
				return true
			}
		}
	}
	// o handler passado nao esta registrado
	return false
}

func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

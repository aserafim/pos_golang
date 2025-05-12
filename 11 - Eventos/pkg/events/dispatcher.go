package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// Cria a struct EventDispatcher
// Classe/struct que implemente
// um EventDispatcher
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

// Cria um construtor de EventDispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Cria a função Dispatch
func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, &wg)
		}
		wg.Wait()
	}

	return nil
}

// Cria a função para registrar eventos
func (ed *EventDispatcher) Register(evName string, hand EventHandlerInterface) error {
	// Associa um handler a um evento
	// Verifica se o handler já está
	// registrado
	_, ok := ed.handlers[evName]
	if ok {
		for _, h := range ed.handlers[evName] {
			if h == hand {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[evName] = append(ed.handlers[evName], hand)

	return nil
}

// Cria função para limpar os handlers
func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

// Cria função para verificar a existência de um
// evento e handler
func (ed *EventDispatcher) Has(evName string, hand EventHandlerInterface) bool {
	_, ok := ed.handlers[evName]
	if ok {
		for _, h := range ed.handlers[evName] {
			if h == hand {
				return true
			}
		}
	}
	return false
}

// Cria a função para remover um handler
func (ed *EventDispatcher) Remove(evName string, hand EventHandlerInterface) error {
	_, ok := ed.handlers[evName]
	if ok {
		for i, h := range ed.handlers[evName] {
			if h == hand {
				ed.handlers[evName] = append(ed.handlers[evName][:i], ed.handlers[evName][i+1:]...)
				return nil
			}
		}
	}

	return nil
}

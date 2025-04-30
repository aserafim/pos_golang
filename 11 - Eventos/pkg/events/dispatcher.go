package events

import "errors"

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

// Cria a função para registrar eventos
func (ed *EventDispatcher) Register(evName string, hand EventHandlerInterface) error {
	// Associa um handler a um evento
	// Verifica se o handler já está
	// registrado
	_, ok := ed.handlers[evName]
	if ok {
		for _, h := range ed.handlers[evName] {
			if h == hand {
				return errors.New("handler already registered")
			}
		}
	}

	ed.handlers[evName] = append(ed.handlers[evName], hand)

	return nil
}

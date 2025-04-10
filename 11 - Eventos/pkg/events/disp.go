/* package events

import "errors"

// Cria a struct EventDispatcher
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

// Instancia um novo EventDispatcher
func NewDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Cria a função para registrar eventos
func (d *EventDispatcher) Register(name string, hand EventHandlerInterface) error {
	_, ok := d.handlers[name]
	if ok {
		for _, h := range d.handlers[name] {
			if h == hand {
				return errors.New("event already exists")
			}
		}
	}
	d.handlers[name] = append(d.handlers[name], hand)
	return nil
} */
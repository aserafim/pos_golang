package events

// Cria a struct EventDispatcher
// Classe/struct que implemente
// uma interface
type EventDispatcher struct {
	handlers map[string][]EventDispatcherInterface
}

// Cria um construtor de EventDispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		make(map[string][]EventDispatcherInterface),
	}
}

// Cria a função para registrar eventos
func (ed *EventDispatcher) Register(evName string, hand EventHandlerInterface) error {

	return nil
}

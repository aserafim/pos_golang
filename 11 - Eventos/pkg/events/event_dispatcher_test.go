package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

// getName
func (e *TestEvent) getName() string {
	return e.Name
}

// getPayload
func (e *TestEvent) getPayload() interface{} {
	return e.Payload
}

// getDateTime
func (e *TestEvent) getDateTime() time.Time {
	return time.Now()
}

// EventHandler
type TestEventHandler struct {
	ID int
}

// HandleFunc
func (h *TestEventHandler) Handle(event EventInterface) {}

// Cria a suite de Dispatcher
type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.event = TestEvent{Name: "Event1", Payload: "test1:1"}
	suite.event2 = TestEvent{Name: "Event2", Payload: "test2:2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.getName(), &suite.handler)
	suite.Nil(err)
	// Comparar 1 com o len da quantidade de handlers
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.getName()]))

	//Registrar outro handler
	err = suite.eventDispatcher.Register(suite.event.getName(), &suite.handler2)
	suite.Nil(err)

	// Validar se foram dois
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.getName()]))

	// Verificar se o handler 1 registrado é o mesmo que passamos
	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.getName()][0])

	// Verificar se o handler 2 registrado é o meso que passamos
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.getName()][1])

}

func (suite *EventDispatcherTestSuite) TestDuplicateEvent() {
	err := suite.eventDispatcher.Register(suite.event.getName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers))

	err = suite.eventDispatcher.Register(suite.event.getName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers))

}

func (suite *EventDispatcherTestSuite) TestClearHandler() {
	// Adiciona um evento e um handler
	err := suite.eventDispatcher.Register(suite.event.getName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.getName()]))

	// Registra mais um handler ao evento
	err = suite.eventDispatcher.Register(suite.event.getName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.getName()]))

	// Limpa os handlers
	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))

}

func (suite *EventDispatcherTestSuite) TestHas() {
	// Adiciona um evento a um handler
	err := suite.eventDispatcher.Register(suite.event.getName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.getName()]))

	// Adiciona mais um evento e handler
	err = suite.eventDispatcher.Register(suite.event.getName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.getName()]))

	// Verifca se existe um handler associado aos eventos
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.getName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.getName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.getName(), &suite.handler3))

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

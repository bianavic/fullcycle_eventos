package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	event1          TestEvent
	event2          TestEvent
	handler1        TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler1 = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event1 = TestEvent{Name: "test1", Payload: "test1"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (suite *EventDispatcherTestSuite) TeardownTest() {
	suite.eventDispatcher = nil
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// se o handler registrado é o mesmo que passamos
	assert.Equal(suite.T(), &suite.handler1, suite.eventDispatcher.handlers[suite.event1.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event1.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// se o handler ja foi registrado para aquele evento
	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Clear() {
	// Event1 com 2 handlers registrados
	// registra evento1 no dispatcher handler1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registra evento1 no dispatcher handler2
	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// Event2 com 1 handler registrado
	// registra evento2 no dispatcher handler3
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	err = suite.eventDispatcher.Clear()
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Has() {
	// Event1 com 2 handlers registrados
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event1.GetName(), &suite.handler1))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event1.GetName(), &suite.handler2))
	// handler 3 nao esta registrado
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event1.GetName(), &suite.handler3))
}

// ao rodar TestSuite, todos os metodos sao executados
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

package interactor

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/repository"
)

func TestSomething(t *testing.T) {
	ctrl := gomock.NewController(t)

	assert := assert.New(t)

	m := repository.NewMockMessageRepository(ctrl)
	m.
		EXPECT().
		Publish("test").
		Return(nil)

	interactor := NewMessageInteractor(m)

	err := interactor.Publish("test")

	// ssert equality
	assert.Equal(err, nil, "they should be equal")

	cb := func(message model.Message) {
		fmt.Printf("[MQTT] Received to %s [Received Message: %s]\n", message.Topic, message.Payload)
	}

	m2 := repository.NewMockMessageRepository(ctrl)
	m2.
		EXPECT().
		Subscribe(gomock.Any()).
		Return(nil)

	interactor2 := NewMessageInteractor(m2)

	err = interactor2.Subscribe(cb)

	assert.Equal(err, nil, "they should be equal")

	// // assert inequality
	// assert.NotEqual(123, 456, "they should not be equal")

	// // assert for nil (good for errors)
	// assert.Nil(object)

	// // assert for not nil (good when you expect something)
	// if assert.NotNil(object) {

	//   // now we know that object isn't nil, we are safe to make
	//   // further assertions without causing any errors
	//   assert.Equal("Something", object.Value)
	// }
}

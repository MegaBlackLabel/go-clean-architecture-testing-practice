package pubsub

import (
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/repository"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/infrastructure/pubsub/mqtt"
)

// ここではinfrastrucureのuserMemoryとの接続を行う。その際にdomain/repositoryのインターフェイスの実装を行う

type messageMqtt struct {
	Mqtt mqtt.PubSub
}

// NewMessageMqtt userRepository構造体のインターフェイス
func NewMessageMqtt(mqtt mqtt.PubSub) repository.MessageRepository {
	return &messageMqtt{Mqtt: mqtt}
}

func (m *messageMqtt) Publish(message string) error {
	return m.Mqtt.Publish(message)
}

func (m *messageMqtt) Subscribe(callback model.Callback) error {
	cb := func(message mqtt.Message) {
		callback(model.Message{
			Topic:   message.Topic,
			Payload: message.Payload,
		})
	}
	return m.Mqtt.Subscribe(cb)
}

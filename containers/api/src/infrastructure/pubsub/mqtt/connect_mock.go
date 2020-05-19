package mqtt

import (
	"bytes"
	"encoding/gob"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/infrastructure/pubsub"
)

// MQTTコネクタのモック。MQTT.Clientを実装してmqtt.Queをモックとして実装する

var clientMock MQTT.Client
var tokenMock MQTT.Token
var messageMock MQTT.Message

func init() {
	clientMock = NewMqttClientMock()
	tokenMock = NewMqttTokenMock()
	messageMock = NewMqttMessageMock()
}

type client struct {
	cb MQTT.MessageHandler
}

type token struct{}

type message struct {
	duplicate bool
	qos       byte
	retained  bool
	topic     string
	messageID uint16
	payload   []byte
	once      sync.Once
	ack       func()
}

// QueMock -.
type QueMock struct {
	Sendtopic string
	Resvtopic string
	Qos       int
	Retained  bool
	Client    MQTT.Client
}

// NewClient -
func NewClient() MQTT.Client {
	return clientMock
}

// NewQueMock -
func NewQueMock() *pubsub.Que {
	return &pubsub.Que{Client: clientMock}
}

// NewMqttTokenMock -
func NewMqttTokenMock() MQTT.Token {
	return &token{}
}

// NewMqttClientMock -
func NewMqttClientMock() MQTT.Client {
	return &client{}
}

// NewMqttMessageMock -
func NewMqttMessageMock() MQTT.Message {
	return &message{}
}

// MQtt.Tokenインターフェイスの実装

func (t *token) Wait() bool {
	return true
}

func (t *token) WaitTimeout(tt time.Duration) bool {
	return true
}

func (t *token) Error() error {
	return nil
}

// MQTT.Clientインターフェイスの実装

func (c *client) IsConnected() bool {
	return true
}

func (c *client) IsConnectionOpen() bool {
	return true
}

func (c *client) Connect() MQTT.Token {
	return tokenMock
}

func (c *client) Disconnect(quiesce uint) {

}

func (c *client) Publish(topic string, qos byte, retained bool, payload interface{}) MQTT.Token {
	b, err := getBytes(payload)
	if err == nil {
		messageMock = &message{payload: b}
		if c.cb != nil {
			c.cb(clientMock, messageMock)
		}
	}
	return tokenMock
}

func (c *client) Subscribe(topic string, qos byte, callback MQTT.MessageHandler) MQTT.Token {
	if callback != nil {
		c.cb = callback
	}
	return tokenMock
}

func (c *client) SubscribeMultiple(filters map[string]byte, callback MQTT.MessageHandler) MQTT.Token {
	return tokenMock
}

func (c *client) Unsubscribe(topics ...string) MQTT.Token {
	return tokenMock
}

func (c *client) AddRoute(topic string, callback MQTT.MessageHandler) {

}

func (c *client) OptionsReader() MQTT.ClientOptionsReader {
	return MQTT.ClientOptionsReader{}
}

// MQTT.Messageインターフェイスの実装

func (m *message) Duplicate() bool {
	return m.duplicate
}

func (m *message) Qos() byte {
	return m.qos
}

func (m *message) Retained() bool {
	return m.retained
}

func (m *message) Topic() string {
	return m.topic
}

func (m *message) MessageID() uint16 {
	return m.messageID
}

func (m *message) Payload() []byte {
	return m.payload
}

func (m *message) Ack() {
	m.once.Do(m.ack)
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

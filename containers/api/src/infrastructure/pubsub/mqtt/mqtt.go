package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/infrastructure/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/utils/logs"
)

// Publish、Subscribeを実装する。SubscribeについてはCallbackをうまくやる

// PubSub インターフェイス
type PubSub interface {
	Publish(message string) error
	Subscribe(callback Callback) error
}

// Message Subscribeのメッセージ
type Message struct {
	Topic   string
	Payload string
}

// Callback コールバック関数定義
type Callback func(message Message)

// Mqtt 構造体
type Mqtt struct {
	Sendtopic string
	Resvtopic string
	Qos       int
	Retained  bool
	Client    MQTT.Client
	logs      logs.Logger
}

// NewMqtt MQTTのクライアント実装 [この関数をwireでDI]
func NewMqtt(que *pubsub.Que, logs logs.Logger) PubSub {
	return &Mqtt{
		Sendtopic: que.Sendtopic,
		Resvtopic: que.Resvtopic,
		Qos:       que.Qos,
		Retained:  que.Retained,
		Client:    que.Client,
		logs:      logs,
	}
}

// Publish MQTTメッセージ送信
func (q *Mqtt) Publish(message string) error {
	if q.Client != nil {
		token := q.Client.Publish(q.Sendtopic, byte(q.Qos), q.Retained, message)
		if token == nil {
			return token.Error()
		}
		q.logs.Infof("[MQTT] Sent to %s\n", q.Sendtopic)
	}
	return nil
}

// Subscribe メッセージ受信コールバック関数設定
func (q *Mqtt) Subscribe(callback Callback) error {
	if q.Client != nil && callback != nil {
		// コールバック
		cb := func(client MQTT.Client, message MQTT.Message) {
			// q.logs.Infof("[MQTT] Received to %s [Received Message: %s]\n", message.Topic(), message.Payload())
			callback(Message{Topic: message.Topic(), Payload: string(message.Payload())})
		}
		if token := q.Client.Subscribe(q.Resvtopic, byte(q.Qos), cb); token.Wait() && token.Error() != nil {
			return token.Error()
		}
		q.logs.Infof("[MQTT] Subscribe to %s\n", q.Resvtopic)
	}

	return nil
}

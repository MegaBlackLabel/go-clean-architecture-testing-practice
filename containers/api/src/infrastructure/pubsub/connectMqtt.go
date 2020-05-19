//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/MegaBlackLabel/go-clean-architecture-testing-practice/infrastructure/pubsub/$GOPACKAGE

package pubsub

import (
	"crypto/tls"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// 接続処理をしてMQTT.Clientを返す

// MqttConfig 接続時の設定項目
type MqttConfig struct {
	Server    string
	Sendtopic string
	Resvtopic string
	Qos       int
	Retained  bool
	Clientid  string
	Username  string
	Password  string
}

// Que -.
type Que struct {
	Sendtopic string
	Resvtopic string
	Qos       int
	Retained  bool
	Client    MQTT.Client
}

// NewMqttConnect MQTT接続
func NewMqttConnect(config MqttConfig) (*Que, error) {
	connOpts := MQTT.NewClientOptions().AddBroker(config.Server).SetClientID(config.Clientid).SetCleanSession(true)
	if config.Username != "" {
		connOpts.SetUsername(config.Username)
		if config.Password != "" {
			connOpts.SetPassword(config.Password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	fmt.Printf("[MQTT] Connected to %s\n", config.Server)

	que := &Que{
		Sendtopic: config.Sendtopic,
		Resvtopic: config.Resvtopic,
		Qos:       config.Qos,
		Retained:  config.Retained,
		Client:    client,
	}

	return que, nil
}

package model

// Message Subscribeのメッセージ
type Message struct {
	Topic   string
	Payload string
}

// Callback コールバック関数定義
type Callback func(message Message)

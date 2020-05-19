package handler

import (
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/adapter/controllers"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/model"
)

// MessageHandler インターフェイス。外部層との接続で使う
type MessageHandler interface {
	Publish(message string) error
	Subscribe(callback model.Callback) error
}

type messageHandler struct {
	MessageController controllers.MessageController
}

// NewMessageHandler resoleverのハンドラー処理。adapter/controllersを呼び出す [この関数をwireでDI]
func NewMessageHandler(mc controllers.MessageController) MessageHandler {
	return &messageHandler{MessageController: mc}
}

// Publish メッセージ送信
func (mh *messageHandler) Publish(message string) error {
	// fmt.Printf("messageHandler %v", message)
	return mh.MessageController.Publish(message)
}

// Subscribe メッセージ受信コールバック関数設定
func (mh *messageHandler) Subscribe(callback model.Callback) error {
	// fmt.Printf("func %v", callback)
	return mh.MessageController.Subscribe(callback)
}

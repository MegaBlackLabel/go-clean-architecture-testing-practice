package controllers

import (
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/usecase/interactor"
)

// ここでは外部から受け取ったデータをusecase/interactorの実装にわたす。コントローラーでは処理の実装をしない
// 受け取ったデータを構造体でusecaseに渡したい場合はusecase側でinputを作って受け渡し用の構造体を用意する
// MQTTの場合、ほぼデータをスルーするのみ

// MessageController インターフェイス
type MessageController interface {
	Publish(message string) error
	Subscribe(callback model.Callback) error
}

type messageController struct {
	MessageInteractor interactor.MessageInteractor
}

// NewMessageController userのコントローラー。処理の実態であるusecase/interactorを呼び出す [この関数をwireでDI]
func NewMessageController(us interactor.MessageInteractor) MessageController {
	return &messageController{MessageInteractor: us}
}

// Publish メッセージ送信
func (mc *messageController) Publish(message string) error {
	return mc.MessageInteractor.Publish(message)
}

// Subscribe メッセージ受信設定
func (mc *messageController) Subscribe(callback model.Callback) error {
	return mc.MessageInteractor.Subscribe(callback)
}

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/MegaBlackLabel/go-clean-architecture-testing-practice/usecase/interactor/$GOPACKAGE

package interactor

import (
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/repository"
)

// ここではmessageInteractorの構造体を定義してusecase/repositoryを実装する

// MessageInteractor  インターフェイスはNewMessageInteractorで使用する
type MessageInteractor interface {
	Publish(message string) error
	Subscribe(callback model.Callback) error
}

type messageInteractor struct {
	MessageRepository repository.MessageRepository
}

// NewMessageInteractor messageの実処理。実処理と登録・参照処理としてusecase/repositoryを呼び出す [この関数をwireでDI]
func NewMessageInteractor(mr repository.MessageRepository) MessageInteractor {
	return &messageInteractor{MessageRepository: mr}
}

// Publish メッセージ送信
func (mi *messageInteractor) Publish(message string) error {
	return mi.MessageRepository.Publish(message)
}

// Subscribe メッセージ受信設定
func (mi *messageInteractor) Subscribe(callback model.Callback) error {
	return mi.MessageRepository.Subscribe(callback)
}

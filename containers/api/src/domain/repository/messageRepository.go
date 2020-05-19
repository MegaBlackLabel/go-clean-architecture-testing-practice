//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/repository/$GOPACKAGE

package repository

import "github.com/MegaBlackLabel/go-clean-architecture-testing-practice/domain/model"

// ここではRepositoryのインターフェイスを定義。adapter/gatewaysで実装する

// MessageRepository インターフェイス wireの関係でusecaseとadapterから参照されるのでdomainにあるほうが最適
type MessageRepository interface {
	Publish(message string) error
	Subscribe(callback model.Callback) error
}

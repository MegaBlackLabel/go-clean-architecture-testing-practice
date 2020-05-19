//+build wireinject

package registry

// The build tag makes sure the stub is not built in the final build.
import (
	"github.com/google/wire"

	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/adapter/controllers"
	pubsubAdapter "github.com/MegaBlackLabel/go-clean-architecture-testing-practice/adapter/gateways/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/infrastructure/handler"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/infrastructure/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/infrastructure/pubsub/mqtt"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/usecase/interactor"
	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/utils/logs/lggr"
)

func InitialiseMqtt(config pubsub.MqttConfig) (handler.MessageHandler, error) {
	wire.Build(
		lggr.NewLogger,
		pubsub.NewMqttConnect,
		mqtt.NewMqtt,
		pubsubAdapter.NewMessageMqtt,
		interactor.NewMessageInteractor,
		controllers.NewMessageController,
		handler.NewMessageHandler,
	)
	return nil, nil
}

func InitialiseMqttMock() mqtt.PubSub {
	wire.Build(
		lggr.NewLogger,
		mqtt.NewQueMock,
		mqtt.NewMqtt,
	)
	return nil
}

//go:build wireinject
// +build wireinject

package wire

import (
	"webhooklenna/config"
	"webhooklenna/controllers"
	webchatrepository "webhooklenna/repositories/webhookRepository/webchatRepository"
	webchatservice "webhooklenna/services/webhookService/webchatService"

	"github.com/google/wire"

	webchatController "webhooklenna/controllers/webhookController/webchatController"
)

var WebhookWebchatSet = wire.NewSet(
	config.Postgres,
	webchatrepository.NewWebhookWebchatRepository,
	wire.Bind(new(webchatrepository.WebhookWebchatRepository), new(*webchatrepository.WebhookWebchatRepositoryImpl)),
	webchatservice.NewWebhookWebchatService,
	wire.Bind(new(webchatservice.WebhookWebchatService), new(*webchatservice.WebhookWebchatServiceImpl)),
	webchatController.NewWebhookWebchatController,
	wire.Bind(new(webchatController.WebhookWebchatController), new(*webchatController.WebhookWebchatControllerImpl)),
	controllers.NewAllController,
)

func InitializeController() *controllers.AllController {
	wire.Build(
		WebhookWebchatSet,
	)
	return nil
}

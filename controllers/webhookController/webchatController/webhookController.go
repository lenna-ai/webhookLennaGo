package webchatcontroller

import (
	webchatservice "webhooklenna/services/webhookService/webchatService"

	"github.com/gofiber/fiber/v2"
)

type WebhookWebchatControllerImpl struct {
	WebhookWebchatService webchatservice.WebhookWebchatService
}

type WebhookWebchatController interface {
	CreateMessage(c *fiber.Ctx) error
}

func NewWebhookWebchatController(webhookWebchatService webchatservice.WebhookWebchatService) *WebhookWebchatControllerImpl {
	return &WebhookWebchatControllerImpl{
		WebhookWebchatService: webhookWebchatService,
	}
}

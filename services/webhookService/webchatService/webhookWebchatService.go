package webchatservice

import (
	webchatrepository "webhooklenna/repositories/webhookRepository/webchatRepository"

	"github.com/gofiber/fiber/v2"
)

type WebhookWebchatService interface {
	WebhookWebchat(c *fiber.Ctx, senderId int, integrationId int, event string, text string) (any, error)
	EncryptText(key string, plaintext string) (string, error)
}

type WebhookWebchatServiceImpl struct {
	webhookWebchatRepository webchatrepository.WebhookWebchatRepository
}

func NewWebhookWebchatService(webhookWebchatRepository webchatrepository.WebhookWebchatRepository) *WebhookWebchatServiceImpl {
	return &WebhookWebchatServiceImpl{
		webhookWebchatRepository: webhookWebchatRepository,
	}
}

package webchatcontroller

import (
	"webhooklenna/helpers"

	"github.com/gofiber/fiber/v2"
)

func (WebhookWebchatControllerImpl *WebhookWebchatControllerImpl) CreateMessage(c *fiber.Ctx) error {
	// senderId = 1524940
	// integrationId = 1098
	// events = "message"
	// app_id = 544
	data, err := WebhookWebchatControllerImpl.WebhookWebchatService.WebhookWebchat(c, 1524940, 1098, "message", "text message")
	if err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	result := fiber.Map{
		"result": data,
	}
	return helpers.ResultSuccessJsonApi(c, result)
}

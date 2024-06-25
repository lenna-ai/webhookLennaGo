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
	// return helpers.ResultSuccessJsonApi(c, data)
	// data, statusCode, err := WebhookWebchatControllerImpl.WebhookWebchatService.HitApi(c)
	// if err != nil {
	// 	panic(err.Error())
	// }

	result := fiber.Map{
		"result": data,
	}
	return helpers.ResultSuccessJsonApi(c, result)
}

func (WebhookWebchatControllerImpl *WebhookWebchatControllerImpl) TestEncode64(c *fiber.Ctx) error {
	// senderId = 1524940
	// integrationId = 1098
	// events = "message"
	// app_id = 544
	WebhookWebchatControllerImpl.WebhookWebchatService.EncryptText("k384y0r4nTRIANGLE", "544:1719197008104@gmail.com")
	data := map[string]interface{}{}
	return helpers.ResultSuccessJsonApi(c, data)
}

package webchatcontroller

import (
	"reflect"
	"webhooklenna/helpers"
	"webhooklenna/helpers/handlers"
	formattermessagemodel "webhooklenna/models/MessageModel/FormatterMessageModel"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (WebhookWebchatControllerImpl *WebhookWebchatControllerImpl) CreateMessage(c *fiber.Ctx) error {
	// senderId = 1524940
	// integrationId = 1098
	// events = "message"
	// app_id = 544
	var request = new(formattermessagemodel.CreateMessageModelFormatterRequest)

	if err := c.BodyParser(request); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	err := handlers.Validate.Struct(*request)
	if err != nil {
		for _, valueError := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(*request).FieldByName(valueError.StructField())
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}
			errorValidate := handlers.StatusResponseErrorValidate(jsonTag, valueError.Tag())
			return helpers.ResultFailedJsonApi(c, nil, errorValidate)
		}
	}

	data, err := WebhookWebchatControllerImpl.WebhookWebchatService.WebhookWebchat(c, 1524940, 1098, "message", "text message")
	if err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	result := fiber.Map{
		"result": data,
	}
	return helpers.ResultSuccessJsonApi(c, result)
}

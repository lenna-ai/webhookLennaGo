package routes

import (
	"webhooklenna/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, allController *controllers.AllController) {

	app.Get("/", allController.WebhookWebchatController.CreateMessage)
	app.Get("/testEncode64", allController.WebhookWebchatController.TestEncode64)

}

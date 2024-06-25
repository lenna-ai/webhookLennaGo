package main

import (
	"webhooklenna/apps"
	"webhooklenna/routes"
	"webhooklenna/wire"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	apps.InitConfigAll(app)
	controller := wire.InitializeController()
	routes.Router(app, controller)

	app.Listen(":3000")
}

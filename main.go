package main

import (
	"webhooklenna/apps"
	"webhooklenna/helpers/handlers"
	"webhooklenna/routes"
	"webhooklenna/wire"

	"github.com/gofiber/fiber/v2"
)

func main() {
	defer handlers.CatchPanicHandler()
	app := fiber.New()
	apps.InitConfigAll(app)
	controller := wire.InitializeController()
	routes.Router(app, controller)

	if err := app.Listen(":3000"); err != nil {
		panic(err.Error())
	}
}

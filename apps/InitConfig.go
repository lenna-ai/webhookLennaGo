package apps

import (
	"webhooklenna/config"
	appconfig "webhooklenna/config/appConfig"

	"github.com/gofiber/fiber/v2"
)

func InitConfigAll(app *fiber.App) error {
	config.Logger(app)
	appconfig.InitEnv()
	config.Postgres()
	return nil
}

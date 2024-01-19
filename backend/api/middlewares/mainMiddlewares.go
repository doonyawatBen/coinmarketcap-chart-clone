package middlewares

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetMainMiddlewares(app *fiber.App) {
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())
}

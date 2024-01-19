package modules

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/handlers"
	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/middlewares"

	"github.com/lodashventure/nlp/model"
)

func Webservice() (*fiber.App, *fiber.Router) {
	app := fiber.New(fiber.Config{
		CaseSensitive: true, // /Foo and /foo are different routes
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var ErrorResponse []model.ErrorResponse
			ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: err.Error()})
			return c.Status(fiber.StatusInternalServerError).JSON(infrastructure.Response(false, ErrorResponse))
		},
	})

	/*
		Applied Middleware.
	*/
	middlewares.SetMainMiddlewares(app)
	app.Use(
		middlewares.LogInit,
		middlewares.LogRestHTTP,
	)

	/*
		Routing
	*/
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	/*
		Routing System
	*/

	v1.Get("/health", handlers.HealthCheckHanlder)

	/*
		Routing Other
	*/
	quotaHandler := handlers.NewQuotaHandler(&v1)
	quotaHandler.NewQuotaRoute()

	return app, &v1
}

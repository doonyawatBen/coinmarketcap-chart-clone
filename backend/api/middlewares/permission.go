package middlewares

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/model"
)

func CheckAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ErrorResponse []model.ErrorResponse

		isAdmin := c.Locals("isAdmin").(bool)
		if !isAdmin {
			ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Permission denide"})
			return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
		}

		return c.Next()
	}
}

package middlewares

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/helpers"
	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/model"
)

func CheckToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ErrorResponse []model.ErrorResponse

		tokenId := c.Get("x-token")
		if tokenId == "" {
			ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Please define x-token in header."})
			return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
		}

		credential := helpers.PrepareCredential()
		var appInfo = credential[tokenId]
		appName := appInfo.AppName
		if appName == "" {
			ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "x-token not have app_name."})
		}

		appCode3 := appInfo.AppCode3
		if appCode3 == "" {
			ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "x-token not have app_code_3."})
		}

		if len(ErrorResponse) > 0 {
			return c.Status(http.StatusUnauthorized).JSON(infrastructure.Response(false, ErrorResponse))
		}

		c.Locals("appName", appName)
		c.Locals("isAdmin", appInfo.IsAdmin)
		c.Locals("tokenId", tokenId)
		c.Locals("ip", c.IP())
		c.Locals("path", c.Path())

		return c.Next()
	}
}

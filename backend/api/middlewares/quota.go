package middlewares

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/model"
	"github.com/lodashventure/nlp/service"
)

func CheckQuota() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ErrorResponse []model.ErrorResponse
		appName := c.Locals("appName").(string)
		quotaService := &service.QuotaService{}

		quota, err := quotaService.GetResultByAppName(appName)
		if err != nil {
			ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Can't connect to database."})
			return c.Status(http.StatusInternalServerError).JSON(infrastructure.Response(false, ErrorResponse))
		}

		if quota.MatchedCount == 1 {
			if quota.ModifiedCount == 0 {
				ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Monthly quota exceeded"})
				return c.Status(http.StatusInternalServerError).JSON(infrastructure.Response(false, ErrorResponse))
			}
		} else {
			err := quotaService.InsertByAppName(appName, infrastructure.ConfigGlobal.Quota.InitPerMonth)
			if err != nil {
				ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: err.Error()})
				return c.Status(http.StatusInternalServerError).JSON(infrastructure.Response(false, ErrorResponse))
			}
		}

		return c.Next()
	}
}

package handlers

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/middlewares"
	"github.com/lodashventure/nlp/model"
	"github.com/lodashventure/nlp/service"
)

type NewQuotaClass struct {
	router       *fiber.Router
	quotaService service.QuotaService
}

func NewQuotaHandler(router *fiber.Router) *NewQuotaClass {
	return &NewQuotaClass{
		router: router,
	}
}

func (property *NewQuotaClass) NewQuotaRoute() {
	(*property.router).Get("/quota", middlewares.CheckToken(), property.GetQuota)
	(*property.router).Put("/quota/:app_name", middlewares.CheckToken(), middlewares.CheckAdmin(), property.UpdateByAppName)
	(*property.router).Patch("/quota/reset/:app_name", middlewares.CheckToken(), middlewares.CheckAdmin(), property.UpdateResetByAppName)
	(*property.router).Patch("/quota/reset_all", middlewares.CheckToken(), middlewares.CheckAdmin(), property.UpdateResetAll)
}

func (property *NewQuotaClass) GetQuota(c *fiber.Ctx) error {
	var ErrorResponse []model.ErrorResponse
	appName := c.Locals("appName").(string)

	quota, err := property.quotaService.GetByAppName(appName)
	if err != nil {
		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: err.Error()})
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	quota.QuotaRemaining = quota.Quota - quota.QuotaUsed

	return c.Status(http.StatusOK).JSON(quota)
}

func (property *NewQuotaClass) UpdateByAppName(c *fiber.Ctx) error {
	var ErrorResponse []model.ErrorResponse
	appName := c.Params("app_name")
	if appName == "" || appName == ":app_name" {
		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "app_name is required"})
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	var quota *model.UpdateQuotaRequest
	err := c.BodyParser(&quota)
	if err != nil {
		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Request query invalid format or try again later."})
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	err = property.quotaService.UpdateByAppName(appName, quota)
	if err != nil {
		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: err.Error()})
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	if len(ErrorResponse) > 0 {
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	return c.Status(http.StatusOK).JSON(&model.UpdateQuotaRes{
		Status:  "Success",
		Quota:   quota.Quota,
		AppName: appName,
	})
}

func (property *NewQuotaClass) UpdateResetByAppName(c *fiber.Ctx) error {
	var ErrorResponse []model.ErrorResponse
	appName := c.Params("app_name")
	if appName == "" || appName == ":app_name" {
		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "app_name is required"})
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	err := property.quotaService.UpdateResetByAppName(appName)
	if err != nil {
		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: err.Error()})
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	if len(ErrorResponse) > 0 {
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	return c.Status(http.StatusOK).JSON(&model.ResetQuotaRes{
		Status:    "Success",
		QuotaUsed: 0,
		AppName:   appName,
	})
}

func (property *NewQuotaClass) UpdateResetAll(c *fiber.Ctx) error {
	var ErrorResponse []model.ErrorResponse

	err := property.quotaService.UpdateResetAll()
	if err != nil {
		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: err.Error()})
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	if len(ErrorResponse) > 0 {
		return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, ErrorResponse))
	}

	return c.Status(http.StatusOK).JSON(&model.ResetAllQuotaRes{
		Status: "Success",
	})
}

package middlewares

import (
	"net/http"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/model"
)

func LogInit(c *fiber.Ctx) error {
	c.Locals("appName", "")
	c.Locals("tokenId", "")
	c.Locals("ip", "")
	c.Locals("path", "")

	return c.Next()
}

func LogRestHTTP(c *fiber.Ctx) error {
	defer func(c *fiber.Ctx) {
		timeNow := time.Now()
		// Use a deferred function to insert the log entry into MongoDB
		logHTTP := model.LogHTTP{
			Method:    c.Method(),
			Status:    c.Response().StatusCode(),
			StatusMsg: http.StatusText(c.Response().StatusCode()),
			Path:      c.Path(),
			IP:        c.IP(),
			AppName:   c.Locals("appName").(string),
			TokenID:   c.Locals("tokenId").(string),
			Duration:  time.Since(timeNow).Milliseconds(),
			Time:      timeNow.UTC().Format("2006-01-02 15:04:05"),
			CreatedAt: timeNow.Format(time.RFC3339),
		}

		if r := recover(); r != nil {
			// Update the response status code to 500 if a panic occurred
			c.Status(fiber.StatusInternalServerError)
			logHTTP.Status = fiber.StatusInternalServerError
			logHTTP.StatusMsg = http.StatusText(fiber.StatusInternalServerError)
		}

		go func(logHTTP *model.LogHTTP) {
			err := infrastructure.LogServiceManage.Insert(logHTTP)
			if err != nil {
				infrastructure.Log.Error(nil, "LogRestHTTP", logHTTP.AppName, "Can't insert data log http in database", err)
			}
		}(&logHTTP)
	}(c)

	return c.Next()
}

package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func RegisterRoute(version *fiber.Router) {
	router := (*version).Group("/graph") // /api/v1/graph

	GraphHandler := NewGraphHandler(&router)
	GraphHandler.NewGraphRoute()
}

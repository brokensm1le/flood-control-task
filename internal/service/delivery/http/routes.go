package http

import "github.com/gofiber/fiber/v2"

func MapRoutes(router fiber.Router, s *ServiceHandler) {
	userGroup := router.Group("/user")

	userGroup.Get("/:username", s.CheckFlood())
}

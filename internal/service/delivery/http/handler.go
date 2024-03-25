package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"task/internal/service"
)

type ServiceHandler struct {
	fcUC service.FloodControl
}

func NewServiceHandler(fcUC service.FloodControl) *ServiceHandler {
	return &ServiceHandler{
		fcUC: fcUC,
	}
}

func (h *ServiceHandler) CheckFlood() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			resp bool
			err  error
		)

		userIdStr := c.Params("username")
		if userIdStr == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Bad request"})
		}

		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Bad request"})
		}
		resp, err = h.fcUC.Check(context.TODO(), userId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(service.ResponseModel{Result: resp, Err: err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(service.ResponseModel{Result: resp})
	}
}

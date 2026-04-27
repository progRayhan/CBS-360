package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/progRayhan/CBS-360/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParm := c.Params("id")

	id, err := strconv.ParseUint(idParm, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user id",
		})
	}

	user, err := h.service.GetUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

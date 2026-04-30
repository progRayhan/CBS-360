package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/progRayhan/CBS-360/internal/api/dto/request"
	"github.com/progRayhan/CBS-360/internal/mapper"
	"github.com/progRayhan/CBS-360/internal/service"
)

type KYCHandler struct {
	service  service.KYCService
	validate *validator.Validate
}

func NewKYCHandler(s service.KYCService) *KYCHandler {
	return &KYCHandler{
		service:  s,
		validate: validator.New(),
	}
}

func (h *KYCHandler) CreateKYC(c *fiber.Ctx) error {
	var req request.CreateKYCRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	kyc := mapper.ToKYCModel(&req)

	if err := h.service.CreateKYC(kyc); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(kyc)
}

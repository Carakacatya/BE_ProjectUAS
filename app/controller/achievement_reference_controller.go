package controller

import (
	"context"
	"projectuasbe/app/service"

	"github.com/gofiber/fiber/v2"
)

type AchievementReferenceController struct {
	service *service.AchievementReferenceService
}

func NewAchievementReferenceController(s *service.AchievementReferenceService) *AchievementReferenceController {
	return &AchievementReferenceController{s}
}

// GET /achievement-ref/:id
func (c *AchievementReferenceController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := c.service.GetByID(context.Background(), id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "reference not found"})
	}

	return ctx.JSON(data)
}

// GET /achievement-ref/student/:studentId
func (c *AchievementReferenceController) GetByStudent(ctx *fiber.Ctx) error {
	studentID := ctx.Params("studentId")

	data, err := c.service.GetByStudent(context.Background(), studentID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(data)
}

// PUT /achievement-ref/status/:id
func (c *AchievementReferenceController) UpdateStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var body struct {
		Status     string  `json:"status"`
		Note       *string `json:"note"`
		VerifiedBy *string `json:"verified_by"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}

	err := c.service.UpdateStatus(context.Background(), id, body.Status, body.Note, body.VerifiedBy)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"status": body.Status})
}

// DELETE /achievement-ref/:id
func (c *AchievementReferenceController) SoftDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.service.SoftDelete(context.Background(), id)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"deleted": true})
}

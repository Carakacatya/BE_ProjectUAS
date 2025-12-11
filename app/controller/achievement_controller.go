package controller

import (
	"context"
	"projectuasbe/app/service"
	"projectuasbe/app/model/MongoDB"
	"github.com/gofiber/fiber/v2"
)

type AchievementController struct {
	service *service.AchievementService
}

func NewAchievementController(s *service.AchievementService) *AchievementController {
	return &AchievementController{s}
}

// POST /achievements
func (c *AchievementController) Create(ctx *fiber.Ctx) error {
	var body model.Achievement
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}

	id, err := c.service.CreateAchievement(context.Background(), &body)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"id": id})
}

func (c *AchievementController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result, err := c.service.GetByID(context.Background(), id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return ctx.JSON(result)
}

func (c *AchievementController) GetByStudent(ctx *fiber.Ctx) error {
	studentID := ctx.Params("studentId")
	list, err := c.service.GetByStudent(context.Background(), studentID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(list)
}

func (c *AchievementController) Submit(ctx *fiber.Ctx) error {
	refID := ctx.Params("refID")
	err := c.service.Submit(context.Background(), refID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"status": "submitted"})
}

func (c *AchievementController) Verify(ctx *fiber.Ctx) error {
	refID := ctx.Params("refID")
	verifiedBy := ctx.Locals("userId").(string)

	err := c.service.Verify(context.Background(), refID, verifiedBy)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"status": "verified"})
}

func (c *AchievementController) Reject(ctx *fiber.Ctx) error {
	refID := ctx.Params("refID")
	verifiedBy := ctx.Locals("userId").(string)

	var body struct {
		Note string `json:"note"`
	}
	ctx.BodyParser(&body)

	err := c.service.Reject(context.Background(), refID, body.Note, verifiedBy)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"status": "rejected"})
}

func (c *AchievementController) SoftDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.service.SoftDelete(context.Background(), id)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"status": "deleted"})
}

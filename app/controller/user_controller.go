package controller

import (
	"projectuasbe/app/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	resp, err := c.userService.Login(body.Username, body.Password)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(resp)
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		Password string `json:"password"`
		RoleID   string `json:"role_id"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	err := c.userService.Register(body.Username, body.Email, body.FullName, body.Password, body.RoleID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "user created"})
}

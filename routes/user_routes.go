package routes

import (
	"projectuasbe/app/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userController *controller.UserController) {
	api := app.Group("/api/users")

	api.Post("/login", userController.Login)
	api.Post("/register", userController.Register)
}

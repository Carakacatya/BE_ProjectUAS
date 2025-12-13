package route

import (
	"database/sql"
	"projectuasbe/helper"
	"projectuasbe/middleware"
	model "projectuasbe/app/model/Postgresql"
	"projectuasbe/app/repository"
	"projectuasbe/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

API := app.Group("/api/v1")

	API.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API v1 Running"})
	})

authRepo := repository.NewAuthRepository(db)
authService := service.NewAuthService(authRepo)
	


       auth := API.Group("/auth")
	auth.Post("/login", authService.LoginEndpoint)
	auth.Post("/refresh", middleware.AuthMiddleware(""), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Refresh endpoint - to be implemented"})
	})
	auth.Post("/logout", middleware.AuthMiddleware(""), authService.LogoutEndpoint)
	auth.Get("/profile", middleware.AuthMiddleware(""), authService.ProfileEndpoint)
package routes

import (
	"projectuasbe/app/repository"
	"projectuasbe/app/service"
	"projectuasbe/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
)

func APIV1Routes(
	app *fiber.App,
	db *pgxpool.Pool,
	mongoColl *mongo.Collection,
) {

	API := app.Group("/api/v1")

	API.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API v1 Running"})
	})

	// =====================
	// AUTH
	// =====================
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)

	auth := API.Group("/auth")
	auth.Post("/login", authService.LoginEndpoint)
	auth.Post("/refresh", middleware.AuthMiddleware(""), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Refresh endpoint - to be implemented"})
	})
	auth.Post("/logout", middleware.AuthMiddleware(""), authService.LogoutEndpoint)
	auth.Get("/profile", middleware.AuthMiddleware(""), authService.ProfileEndpoint)

	// =====================
	// USERS (ADMIN ONLY)
	// =====================
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	users := API.Group("/users")
	users.Use(middleware.AuthMiddleware("manage_users"))
	users.Get("/", userService.GetUsersEndpoint)
	users.Get("/:id", userService.GetUserByIDEndpoint)
	users.Post("/", userService.CreateUserEndpoint)
	users.Put("/:id", userService.UpdateUserEndpoint)
	users.Delete("/:id", userService.DeleteUserEndpoint)
	users.Put("/:id/role", userService.UpdateUserRoleEndpoint)

	// =====================
	// ACHIEVEMENTS
	// =====================
	achievementRepo := repository.NewAchievementRepository(db, mongoColl)
	achievementService := service.NewAchievementService(achievementRepo)

	achievements := API.Group("/achievements")
	achievements.Use(middleware.AuthMiddleware(""))
	achievements.Get("/", achievementService.GetAchievementsEndpoint)
	achievements.Get("/:id", achievementService.GetAchievementByIDEndpoint)
	achievements.Post("/", achievementService.CreateAchievementEndpoint)
	achievements.Put("/:id", achievementService.UpdateAchievementEndpoint)
	achievements.Delete("/:id", achievementService.DeleteAchievementEndpoint)
	achievements.Post("/:id/submit", achievementService.SubmitAchievementEndpoint)
	achievements.Post("/:id/verify", achievementService.VerifyAchievementEndpoint)
	achievements.Post("/:id/reject", achievementService.RejectAchievementEndpoint)
	achievements.Get("/:id/history", achievementService.GetAchievementHistoryEndpoint)
	achievements.Post("/:id/attachments", achievementService.UploadAttachmentEndpoint)
	achievements.Get("/statistics", achievementService.GetAchievementStatisticsEndpoint)

	// =====================
	// STUDENTS
	// =====================
	students := API.Group("/students")
	students.Use(middleware.AuthMiddleware(""))
	students.Get("/", userService.GetStudentsEndpoint)
	students.Get("/:id", userService.GetStudentByIDEndpoint)
	students.Get("/:id/achievements", userService.GetStudentAchievementsEndpoint)
	students.Put("/:id/advisor",
		middleware.AuthMiddleware("manage_users"),
		userService.UpdateStudentAdvisorEndpoint,
	)

	// =====================
	// LECTURERS
	// =====================
	lecturers := API.Group("/lecturers")
	lecturers.Use(middleware.AuthMiddleware(""))
	lecturers.Get("/", userService.GetLecturersEndpoint)
	lecturers.Get("/:id/advisees", userService.GetLecturerAdviseesEndpoint)
}

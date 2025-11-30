package main

import (
	"log"
	"projectuasbe/app/controller"
	"projectuasbe/app/repository"
	"projectuasbe/app/service"
	"projectuasbe/config"
	"projectuasbe/database"
	"projectuasbe/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load .env
	config.LoadConfig()

	// Init Databases
	database.InitPostgres()
	database.InitMongo()

	app := fiber.New()

	// ===========================
	// REPOSITORIES
	// ===========================
	userRepo := repository.NewUserRepository(database.Postgres)
	studentRepo := repository.NewStudentRepository(database.Postgres)
	lecturerRepo := repository.NewLecturerRepository(database.Postgres)
	achievementRefRepo := repository.NewAchievementReferenceRepository(database.Postgres)

	// ===========================
	// SERVICES
	// ===========================
	userService := service.NewUserService(userRepo)
	studentService := service.NewStudentService(studentRepo)
	lecturerService := service.NewLecturerService(lecturerRepo)
	achievementRefService := service.NewAchievementReferenceService(achievementRefRepo)

	// ===========================
	// CONTROLLERS
	// ===========================
	userController := controller.NewUserController(userService)
	studentController := controller.NewStudentController(studentService)
	lecturerController := controller.NewLecturerController(lecturerService)
	achievementRefController := controller.NewAchievementReferenceController(achievementRefService)

	// ===========================
	// ROUTES
	// ===========================
	app.Group("/api")

	routes.UserRoutes(app, userController)
	routes.StudentRoutes(app, studentController)
	routes.LecturerRoutes(app, lecturerController)
	routes.AchievementReferenceRoutes(app, achievementRefController)

	log.Println("🚀 Server running on port", config.AppConfig.AppPort)
	app.Listen(":" + config.AppConfig.AppPort)
}

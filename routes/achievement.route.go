package routes

import (
	"projectuasbe/app/controller"
	"projectuasbe/app/repository"
	"projectuasbe/app/service"
	"projectuasbe/database"
	"projectuasbe/middleware"

	"github.com/gofiber/fiber/v2"
)

func AchievementRoutes(app fiber.Router) {

	// ===== REPOSITORIES =====
	achRepo := repository.NewAchievementRepository()                         // MongoDB
	refRepo := repository.NewAchievementReferenceRepository(database.Postgres) // PostgreSQL

	// ===== SERVICES =====
	achService := service.NewAchievementService(achRepo, refRepo)

	// ===== CONTROLLER =====
	achController := controller.NewAchievementController(achService)

	// ===== ROUTES =====
	r := app.Group("/achievements")

	// Mahasiswa membuat dan membaca prestasi
	r.Post("/", middleware.JWTProtected(), middleware.MahasiswaOnly(), achController.Create)
	r.Get("/:id", middleware.JWTProtected(), middleware.MahasiswaOnly(), achController.GetByID)
	r.Get("/student/:studentId", middleware.JWTProtected(), middleware.MahasiswaOnly(), achController.GetByStudent)

	// Mahasiswa submit prestasi
	r.Put("/submit/:refID", middleware.JWTProtected(), middleware.MahasiswaOnly(), achController.Submit)

	// Dosen verifikasi prestasi
	r.Put("/verify/:refID", middleware.JWTProtected(), middleware.DosenOnly(), achController.Verify)
	r.Put("/reject/:refID", middleware.JWTProtected(), middleware.DosenOnly(), achController.Reject)

	// Mahasiswa hapus prestasi (soft delete)
	r.Delete("/:id", middleware.JWTProtected(), middleware.MahasiswaOnly(), achController.SoftDelete)
}

package routes

import (
	"projectuasbe/app/controller"
	"projectuasbe/middleware"

	"github.com/gofiber/fiber/v2"
)

func AchievementReferenceRoutes(app fiber.Router, c *controller.AchievementReferenceController) {

	r := app.Group("/achievement-ref")

	r.Get("/:id",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "dosen_wali", "mahasiswa"),
		c.GetByID,
	)

	r.Get("/student/:studentId",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "dosen_wali", "mahasiswa"),
		c.GetByStudent,
	)

	r.Put("/status/:id",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "dosen_wali"),
		c.UpdateStatus,
	)

	r.Delete("/:id",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin"),
		c.SoftDelete,
	)
}

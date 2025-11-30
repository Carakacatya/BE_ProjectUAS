package routes

import (
	"projectuasbe/app/controller"
	"projectuasbe/middleware"

	"github.com/gofiber/fiber/v2"
)

func LecturerRoutes(app fiber.Router, lecturerController *controller.LecturerController) {

	r := app.Group("/lecturers")

	r.Get("/:id",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "dosen_wali"),
		lecturerController.GetProfile,
	)

	r.Get("/:id/students",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "dosen_wali"),
		lecturerController.GetStudents,
	)
}

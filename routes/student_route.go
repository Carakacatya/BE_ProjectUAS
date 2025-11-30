package routes

import (
	"projectuasbe/app/controller"
	"projectuasbe/middleware"

	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(app *fiber.App, studentController *controller.StudentController) {
	route := app.Group("/students")

	route.Get("/:id",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "mahasiswa", "dosen_wali"),
		studentController.GetByID)

	route.Get("/user/:userId",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "mahasiswa", "dosen_wali"),
		studentController.GetByUserID)

	route.Get("/advisor/:lecturerId",
		middleware.JWTProtected(),
		middleware.RoleRequired("admin", "dosen_wali"),
		studentController.GetByAdvisorID)
}

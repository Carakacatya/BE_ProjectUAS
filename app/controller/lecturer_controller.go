package controller

import (
	"context"
	"net/http"
	"projectuasbe/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LecturerController struct {
	LecturerService service.LecturerService
}

func NewLecturerController(lecturerService service.LecturerService) *LecturerController {
	return &LecturerController{
		LecturerService: lecturerService,
	}
}

//
// GET PROFILE DOSEN WALI
//
func (lc *LecturerController) GetProfile(c *fiber.Ctx) error {
	idParam := c.Params("id")

	lecturerID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid lecturer ID",
		})
	}

	result, err := lc.LecturerService.GetProfile(context.Background(), lecturerID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

//
// GET MAHASISWA BIMBINGAN
//
func (lc *LecturerController) GetStudents(c *fiber.Ctx) error {
	idParam := c.Params("id")

	lecturerID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid lecturer ID",
		})
	}

	students, err := lc.LecturerService.GetStudents(context.Background(), lecturerID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(students)
}

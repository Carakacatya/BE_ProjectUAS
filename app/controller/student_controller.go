package controller

import (
	"github.com/gofiber/fiber/v2"
	"projectuasbe/app/service"
)

type StudentController struct {
	studentService service.StudentService
}

func NewStudentController(service service.StudentService) *StudentController {
	return &StudentController{
		studentService: service,
	}
}

// =======================================
// Get Student by ID
// GET /students/:id
// =======================================
func (c *StudentController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "student ID is required",
		})
	}

	student, err := c.studentService.GetStudentByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if student == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "student not found",
		})
	}

	return ctx.JSON(student)
}

// =======================================
// Get Student by user_id
// GET /students/user/:userId
// =======================================
func (c *StudentController) GetByUserID(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	if userId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user ID is required",
		})
	}

	student, err := c.studentService.GetStudentByUserID(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if student == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "student not found",
		})
	}

	return ctx.JSON(student)
}

// =======================================
// Get all students under advisor
// GET /students/advisor/:lecturerId
// =======================================
func (c *StudentController) GetByAdvisorID(ctx *fiber.Ctx) error {
	lecturerId := ctx.Params("lecturerId")
	if lecturerId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "lecturer ID is required",
		})
	}

	students, err := c.studentService.GetStudentsByAdvisor(lecturerId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(students)
}

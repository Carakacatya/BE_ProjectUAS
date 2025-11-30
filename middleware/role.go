package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// FLEXIBLE — bisa dipakai untuk banyak role sekaligus
func RoleRequired(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role == nil {
			return c.Status(403).JSON(fiber.Map{
				"error": "Unauthorized: role not found",
			})
		}

		roleStr := role.(string)

		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				return c.Next() // allowed
			}
		}

		return c.Status(403).JSON(fiber.Map{
			"error": "Forbidden: insufficient role permissions",
		})
	}
}

// Shortcut untuk ADMIN
func AdminOnly() fiber.Handler {
	return RoleRequired("admin")
}

// Shortcut untuk DOSEN WALI
func DosenOnly() fiber.Handler {
	return RoleRequired("dosen_wali")
}

// Shortcut untuk MAHASISWA
func MahasiswaOnly() fiber.Handler {
	return RoleRequired("mahasiswa")
}

package response

import "github.com/gofiber/fiber/v2"

type Error struct {
	Message string `json:"message"`
}

func JSON(c *fiber.Ctx, code int, payload any) error {
	return c.Status(code).JSON(payload)
}
func ErrorJSON(c *fiber.Ctx, code int, msg string) error {
	return JSON(c, code, Error{Message: msg})
}

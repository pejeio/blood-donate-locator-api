package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type JSONErrorResponse struct {
	Message string `json:"message"`
}

// GetBearerTokenFromHeaders retrieves the bearer token from the headers of a fiber.Ctx object.
func GetBearerTokenFromHeaders(c *fiber.Ctx) string {
	if c == nil {
		return ""
	}

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func UnauthorizedJSONErrorResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(JSONErrorResponse{
		Message: "Unauthorized",
	})
}

func ForbiddenJSONErrorResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(JSONErrorResponse{
		Message: "Forbidden",
	})
}

func SetUserIDOnCtx(c *fiber.Ctx, value string) {
	c.Locals("userId", value)
}

func GetUserIDFromCtx(c *fiber.Ctx) string {
	return c.Locals("userId").(string)
}

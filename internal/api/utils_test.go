package api

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func TestGetBearerTokenFromHeaders(t *testing.T) {
	t.Parallel()
	app := fiber.New()

	// Test with nil context
	t.Run("Nil Context", func(t *testing.T) {
		token := GetBearerTokenFromHeaders(nil)
		if token != "" {
			t.Errorf("Expected empty token for nil context, but got %s", token)
		}
	})

	// Test with missing Authorization header
	t.Run("Missing Authorization Header", func(t *testing.T) {
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		token := GetBearerTokenFromHeaders(ctx)
		if token != "" {
			t.Errorf("Expected empty token for missing Authorization header, but got %s", token)
		}
	})

	// Test with improperly formatted Authorization header
	t.Run("Improperly Formatted Authorization Header", func(t *testing.T) {
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// ctx.Set("Authorization", "InvalidHeaderFormat")
		ctx.Request().Header.Set("Authorization", "InvalidHeaderFormat")
		token := GetBearerTokenFromHeaders(ctx)
		if token != "" {
			t.Errorf("Expected empty token for improperly formatted Authorization header, but got %s", token)
		}
	})

	// Test with valid Authorization header
	t.Run("Valid Authorization Header", func(t *testing.T) {
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		ctx.Request().Header.Set("Authorization", "Bearer YourValidTokenHere")
		token := GetBearerTokenFromHeaders(ctx)
		expectedToken := "YourValidTokenHere"
		if token != expectedToken {
			t.Errorf("Expected token: %s, but got: %s", expectedToken, token)
		}
	})
}

package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pejeio/blood-donate-locator-api/internal/auth"
	"github.com/rs/zerolog/log"
)

func CorsHandler() fiber.Handler {
	return cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

// Check if the user has rights to create a location
func CreateLocationAuthzMiddleware(authClient *auth.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAllowed := authClient.CheckScopesAllowedOnResource([]string{"create"}, "location", GetBearerTokenFromHeaders(c))
		if !isAllowed {
			return ForbiddenJSONErrorResponse(c)
		}
		return c.Next()
	}
}

// Check if the user has rights to delete a location
func DeleteLocationAuthzMiddleware(authClient *auth.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAllowed := authClient.CheckScopesAllowedOnResource([]string{"delete"}, "location", GetBearerTokenFromHeaders(c))
		if !isAllowed {
			return ForbiddenJSONErrorResponse(c)
		}
		return c.Next()
	}
}

// Just a general middleware that checks for a valid JWT token
func Protect(authClient *auth.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := GetBearerTokenFromHeaders(c)

		log.Debug().Msg(("Retrospecting token..."))
		rptResult, err := authClient.GC.RetrospectToken(c.Context(), accessToken, authClient.ClientID, authClient.ClientSecret, authClient.Realm)
		if err != nil {
			log.Debug().Msg("Token is invalid.")
			return UnauthorizedJSONErrorResponse(c)
		}

		isTokenValid := *rptResult.Active
		if !isTokenValid {
			log.Debug().Msg("Token is invalid.")
			return c.Status(fiber.StatusUnauthorized).JSON(JSONErrorResponse{
				Message: "Unauthorized",
			})
		}

		log.Trace().Msg("Decoding access token...")
		_, claims, err := authClient.GC.DecodeAccessToken(c.Context(), accessToken, authClient.Realm)
		if err != nil {
			log.Debug().Msg("Token is invalid.")
			return UnauthorizedJSONErrorResponse(c)
		}

		if subClaim, ok := (*claims)["sub"].(string); ok {
			SetUserIDOnCtx(c, subClaim)
		}

		return c.Next()
	}
}

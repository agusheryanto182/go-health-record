package middleware

import (
	"strings"

	"github.com/agusheryanto182/go-health-record/module/feature/user"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/gofiber/fiber/v2"
)

func Protected(jwtService jwt.JWTInterface, userService user.UserSvcInterface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			return response.NewUnauthorizedError("Access denied: invalid token")
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		payload, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return response.NewUnauthorizedError("Access denied: invalid token")
		}

		user, err := userService.CheckUserByIdAndRole(payload.Id, payload.Role)
		if err != nil {
			return err
		}

		if !user {
			return response.NewUnauthorizedError("Access denied: invalid token")
		}

		c.Locals("CurrentUser", payload)

		return c.Next()
	}
}

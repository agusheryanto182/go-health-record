package middleware

import (
	"strings"

	"github.com/agusheryanto182/go-health-record/module/feature/user"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
			logrus.Errorf("Error validating token: %v", err)
			return response.NewUnauthorizedError("Access denied: invalid token")
		}

		user, err := userService.GetUserByID(payload.Id)
		if err != nil {
			logrus.Errorf("Error retrieving user: %v", err)
			return response.NewUnauthorizedError("Access denied: user not found")
		}

		c.Locals("CurrentUser", user)

		return c.Next()
	}
}

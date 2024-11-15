package middleware

import (
	"micro/config"
	"micro/internal/models/entity"
	"micro/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	userID := uint(claims["id"].(float64))
	var user entity.Users
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}
// not using feature
	// if !user.Verify {
	// 	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "Account not verified. Please check your email for verification instructions.",
	// 	})
	// }

	c.Locals("usersInfo", claims)
	c.Locals("role", claims["role"])
	return c.Next()
}

func AdminRole(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role == "member" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	return c.Next()
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

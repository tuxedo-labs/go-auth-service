package handler

import (
	"context"
	"errors"
	"fmt"
	"micro/internal/models/request"
	"micro/internal/service"
	"micro/pkg/provider"
	"micro/pkg/utils"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	if errValidate := service.ValidateLogin(loginRequest); errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
	}

	user, err := service.AuthenticateUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	if !user.Verify {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Account not verified. Please check your email for verification instructions.",
		})
	}

	token, errGenerateToken := service.GenerateJWTToken(user)
	if errGenerateToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating token",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"token":  token,
	})
}

func Register(c *fiber.Ctx) error {
	registerRequest := new(request.RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return err
	}

	if errValidate := service.ValidateRegister(registerRequest); errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
	}

	result, err := service.HashAndStoreUser(registerRequest)
	if err != nil {
		if err.Error() == fmt.Sprintf("user with email %s already exists", registerRequest.Email) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Email already in use",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	return c.JSON(fiber.Map{
		"status":  result,
		"message": "Registration successful! Please check your email for the verification code",
	})
}

// Oauth google provider
func AuthGoogle(c *fiber.Ctx) error {
	form := c.Query("from", "/")
	url := service.GetGoogleAuthURL(form)
	return c.Redirect(url)
}

func CallbackAuthGoogle(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Authorization code is missing",
		})
	}

	token, err := provider.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to exchange authorization code for token",
		})
	}

	userInfo, err := service.GetGoogleUserInfo(token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Failed to get user info: %v", err),
		})
	}

	email, emailExists := userInfo["email"].(string)
	givenName := userInfo["given_name"].(string)
	familyName := userInfo["family_name"].(string)

	if !emailExists {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Email is missing from user info",
		})
	}

	existingUser, err := service.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if saveErr := service.SaveGoogleUser(givenName, familyName, email); saveErr != nil {
				return c.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": fmt.Sprintf("Failed to save new user data: %v", saveErr),
				})
			}
			existingUser, err = service.GetUserByEmail(email)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": "Failed to fetch the newly created user",
				})
			}
			jwtToken, err := service.GenerateJWTToken(existingUser)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": "Failed to generate JWT token",
				})
			}

			return c.JSON(fiber.Map{
				"status":  "success",
				"token":   jwtToken,
				"message": "Registered with Google successfully",
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("Failed to check if user exists: %v", err),
			})
		}
	}

	if existingUser.Provider != nil && *existingUser.Provider != "google" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Your account is already registered with provider '%s'", *existingUser.Provider),
		})
	}

	jwtToken, err := service.GenerateJWTToken(existingUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to generate JWT token",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User already exists",
		"token":   jwtToken,
		"data": fiber.Map{
			"user": request.UserResponse{
				ID:        existingUser.ID,
				Name:      existingUser.Name,
				FirstName: existingUser.FirstName,
				LastName:  existingUser.LastName,
				Email:     existingUser.Email,
				CreatedAt: existingUser.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt: existingUser.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			},
		},
	})
}

// GITHU PROVIDER

func AuthGithub(c *fiber.Ctx) error {
	form := c.Query("from", "/")
	url := service.GetGithubAuthUrl(form)
	return c.Redirect(url)
}

func CallbackAuthGithub(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Authorization code is missing",
		})
	}

	token, err := provider.GithubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Error exchanging code for token: %v\n", err)
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to exchange authorization code for token",
		})
	}

	userInfo, err := service.GetGithubUserInfo(token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Failed to get user info: %v", err),
		})
	}

	email, err := service.GetGithubUserPrimaryEmail(token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Failed to get user email: %v", err),
		})
	}

	existingUser, err := service.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			var firstName, lastName string
			if name, ok := userInfo["name"].(string); ok {
				nameParts := strings.Fields(name)
				if len(nameParts) > 0 {
					firstName = nameParts[0]
					if len(nameParts) > 1 {
						lastName = strings.Join(nameParts[1:], " ")
					}
				}
			}
			if saveErr := service.SaveGithubUser(firstName, lastName, email); saveErr != nil {
				return c.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": fmt.Sprintf("Failed to save new user data: %v", saveErr),
				})
			}
			existingUser, err = service.GetUserByEmail(email)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": "Failed to fetch the newly created user",
				})
			}
			jwtToken, err := service.GenerateJWTToken(existingUser)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": "Failed to generate JWT token",
				})
			}

			return c.JSON(fiber.Map{
				"status":  "success",
				"token":   jwtToken,
				"message": "Registered with Github successfully",
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("Failed to check if user exists: %v", err),
			})
		}
	}

	if existingUser.Provider != nil && *existingUser.Provider != "github" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Your account is already registered with provider '%s'", *existingUser.Provider),
		})
	}

	jwtToken, err := service.GenerateJWTToken(existingUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to generate JWT token",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User already exists",
		"token":   jwtToken,
		"data": fiber.Map{
			"user": request.UserResponse{
				ID:        existingUser.ID,
				Name:      existingUser.Name,
				FirstName: existingUser.FirstName,
				LastName:  existingUser.LastName,
				Email:     existingUser.Email,
				CreatedAt: existingUser.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt: existingUser.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			},
		},
	})
}


func VerifyToken(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Unauthorized: Token is missing",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid Token",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Token is valid",
		"claims":  claims,
	})
}

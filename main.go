package main

import (
	"log"
	"micro/config"
	"micro/internal/handler"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.DBConnect()

    // init the router on here
    app.Post("/api/auth/login", handler.Login)
	app.Post("/api/auth/register", handler.Register)

	// oauth google provider
	app.Get("/api/auth/google", handler.AuthGoogle)
	app.Get("/api/auth/google/callback", handler.CallbackAuthGoogle)

	// oauth github provider
	app.Get("/api/auth/github", handler.AuthGithub)
	app.Get("/api/auth/github/callback", handler.CallbackAuthGithub)


	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

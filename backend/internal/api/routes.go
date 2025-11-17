package api

import (
	"github.com/Shihasz/gophiway/internal/config"
	"github.com/Shihasz/gophiway/internal/middleware"
	"github.com/Shihasz/gophiway/internal/repository"
	"github.com/Shihasz/gophiway/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	// API version group
	api := app.Group("/api/" + cfg.APIVersion)

	// Welcome route
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Gophiway API",
			"version": cfg.APIVersion,
			"docs":    "/api/" + cfg.APIVersion + "/docs",
		})
	})

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg)

	// Initialize handlers
	authHandler := NewAuthHandler(authService)

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.RefreshToken)
	auth.Post("/logout", authHandler.Logout)

	// Protected auth routes
	authProtected := api.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware(cfg))
	authProtected.Get("/me", authHandler.GetMe)

	// TODO: Add more route groups here
	// products := api.Group("/products")
	// cart := api.Group("/cart")
	// orders := api.Group("/orders")
}

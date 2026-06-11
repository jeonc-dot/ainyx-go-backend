package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jeon-prince/ainyx-go-backend/internal/handler"
	"github.com/jeon-prince/ainyx-go-backend/internal/middleware"
	"github.com/jeon-prince/ainyx-go-backend/internal/repository"
	"github.com/jeon-prince/ainyx-go-backend/internal/routes"
	"go.uber.org/zap"
)

func main() {
	// 1. Initialize Zap Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 2. Initialize the Validator package
	v := validator.New()

	// 3. Initialize Database connection (Mocked temporarily until Postgres is spun up)
	// dbConn := database.Connect()
	// queries := sqlc.New(dbConn)
	repo := repository.NewUserRepository(nil) // Passing nil just to compile for now

	// 4. Initialize Handlers with injected dependencies
	userHandler := handler.NewUserHandler(repo, v, logger)

	// 5. Initialize the GoFiber App (Like express())
	app := fiber.New()

	// 6. Setup Middlewares (Bonus Requirements: RequestID & Duration Logging)
	app.Use(requestid.New())
	app.Use(middleware.RequestLogger(logger))

	// 7. Setup API Routes
	routes.SetupRoutes(app, userHandler)

	// 8. Start the Server
	logger.Info("Starting server on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

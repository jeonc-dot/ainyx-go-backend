package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jeon-prince/ainyx-go-backend/db/sqlc"
	"github.com/jeon-prince/ainyx-go-backend/internal/models"
	"github.com/jeon-prince/ainyx-go-backend/internal/repository"
	"github.com/jeon-prince/ainyx-go-backend/internal/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	repo     repository.UserRepository
	validate *validator.Validate
	logger   *zap.Logger
}

func NewUserHandler(repo repository.UserRepository, validate *validator.Validate, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		repo:     repo,
		validate: validate,
		logger:   logger,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON payload"})
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Warn("Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed: " + err.Error()})
	}

	dobTime, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format, use YYYY-MM-DD"})
	}

	arg := sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  dobTime,
	}

	user, err := h.repo.CreateUser(c.Context(), arg)
	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.repo.GetUser(c.Context(), int32(id))
	if err != nil {
		h.logger.Error("User not found", zap.Int("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	dobStr := user.Dob.Format("2006-01-02")
	age, err := service.CalculateAge(dobStr)
	if err != nil {
		h.logger.Error("Failed to calculate age", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  dobStr,
		Age:  age,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := h.repo.DeleteUser(c.Context(), int32(id)); err != nil {
		h.logger.Error("Failed to delete user", zap.Int("id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

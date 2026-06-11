package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// RequestLogger intercepts HTTP requests, records the exact duration, 
// and logs the result using the structured Uber Zap logger.
func RequestLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Pass control to the next middleware or handler (like next() in Express)
		err := c.Next()

		duration := time.Since(start)

		logger.Info("HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("request_id", c.GetRespHeader("X-Request-Id")),
		)

		return err
	}
}

package internal

import (
	"discord/app/gateway/internal/router"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func NewHTTPServer(logger *zap.Logger, r router.Router) *fiber.App {
	engine := fiber.New(fiber.Config{
		ErrorHandler: fiber.DefaultErrorHandler,
	})

	engine.Use(cors.New())
	engine.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	engine.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	r.Register(engine)

	return engine
}

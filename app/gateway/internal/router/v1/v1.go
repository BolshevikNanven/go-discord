package v1

import (
	"discord/app/gateway/internal/handler"
	"discord/app/gateway/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	authHandler    *handler.AuthHandler
	authMiddleware *middleware.AuthMiddleware
	bizHandler     *handler.BizHandler
	imHandler      *handler.ImHandler
}

func NewRouter(
	authHandler *handler.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
	bizHandler *handler.BizHandler,
	imHandler *handler.ImHandler,
) *Router {
	return &Router{
		authHandler:    authHandler,
		authMiddleware: authMiddleware,
		bizHandler:     bizHandler,
		imHandler:      imHandler,
	}
}

func (r *Router) Register(engine *fiber.App) {
	v1 := engine.Group("/v1")

	auth := v1.Group("/auth")
	{
		auth.Post("/login", r.authHandler.Login)
		auth.Post("/register", r.authHandler.Register)
		auth.Post("/refresh", r.authHandler.Refresh)
	}

	space := v1.Group("/space").Use(r.authMiddleware.HeaderAccessToken)
	{
		space.Post("", r.bizHandler.CreateSpace)
		space.Post("/join/:space_id", r.bizHandler.JoinSpace)
		space.Post("/leave/:space_id", r.bizHandler.LeaveSpace)
		space.Get("", r.bizHandler.GetSpaces)
		space.Put("/:space_id", r.bizHandler.UpdateSpace)
		space.Delete("/:space_id", r.bizHandler.DeleteSpace)
	}

	channel := v1.Group("/channel").Use(r.authMiddleware.HeaderAccessToken)
	{
		channel.Post("", r.bizHandler.CreateChannel)
		channel.Post("/join/:channel_id", r.bizHandler.JoinChannel)
		channel.Post("/leave/:channel_id", r.bizHandler.LeaveChannel)
		channel.Get("", r.bizHandler.GetChannels)
		channel.Put("/:channel_id", r.bizHandler.UpdateChannel)
		channel.Delete("/:channel_id", r.bizHandler.DeleteChannel)
	}

	message := v1.Group("/message").Use(r.authMiddleware.HeaderAccessToken)
	{
		message.Post("", r.imHandler.SendMessage)
		message.Post("/ack", r.imHandler.AckMessage)
		message.Get("", r.imHandler.GetHistory)
		message.Get("/inbox", r.imHandler.GetInboxMessages)

		message.Post("/channel/:channel_id", r.imHandler.SendMessage)
		message.Post("/channel/:channel_id/ack", r.imHandler.AckChannelMessage)
		message.Get("/channel/:channel_id", r.imHandler.GetHistory)
		message.Get("/channel/:channel_id/inbox", r.imHandler.GetChannelInbox)
	}
}

package middleware

import (
	"discord/app/gateway/internal/view"
	"discord/pkg/jwtutil"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	jwtConfig *jwtutil.Config
}

func NewAuthMiddleware(jwtConfig *jwtutil.Config) *AuthMiddleware {
	return &AuthMiddleware{
		jwtConfig: jwtConfig,
	}
}

func (authMiddleware *AuthMiddleware) HeaderAccessToken(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(view.Response{
			Code:    fiber.StatusUnauthorized,
			Message: "未授权",
			Data:    nil,
		})
	}

	claims, err := jwtutil.ValidateToken(token, jwtutil.AccessToken, authMiddleware.jwtConfig)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(view.Response{
			Code:    fiber.StatusUnauthorized,
			Message: "非法令牌",
			Data:    err.Error(),
		})
	}

	ctx.Locals("user_id", claims.UserId)

	return ctx.Next()
}

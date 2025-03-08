package handler

import (
	"context"
	"discord/api/auth"
	"discord/app/gateway/internal/client"
	"discord/app/gateway/internal/view"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authClientPool *client.AuthClientPool
}

func NewAuthHandler(authClientPool *client.AuthClientPool) *AuthHandler {
	return &AuthHandler{
		authClientPool: authClientPool,
	}
}

func (authHandler *AuthHandler) Login(ctx *fiber.Ctx) error {
	var req view.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    fiber.StatusBadRequest,
			Message: "无效的请求体格式",
			Data:    nil,
		})
	}

	if req.Username == "" || req.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    fiber.StatusBadRequest,
			Message: "用户名和密码不能为空",
			Data:    nil,
		})
	}

	tokenResp, err := authHandler.authClientPool.Get().Login(context.Background(), &auth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		sta, ok := status.FromError(err)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Code:    fiber.StatusInternalServerError,
				Message: "服务器内部错误",
				Data:    nil,
			})
		}

		switch sta.Code() {
		case codes.NotFound, codes.InvalidArgument, codes.AlreadyExists:
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Code:    fiber.StatusBadRequest,
				Message: sta.Message(),
				Data:    nil,
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Code:    fiber.StatusInternalServerError,
				Message: "服务器内部错误",
				Data:    nil,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(view.Response{
		Code:    fiber.StatusOK,
		Message: "登录成功",
		Data: view.TokenResponse{
			AccessToken:  tokenResp.AccessToken,
			RefreshToken: tokenResp.RefreshToken,
		},
	})
}

func (authHandler *AuthHandler) Register(ctx *fiber.Ctx) error {
	var req view.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    fiber.StatusBadRequest,
			Message: "无效的请求体格式",
			Data:    nil,
		})
	}

	if req.Username == "" || req.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    fiber.StatusBadRequest,
			Message: "用户名和密码不能为空",
			Data:    nil,
		})
	}

	user, err := authHandler.authClientPool.Get().Register(context.Background(), &auth.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		sta, ok := status.FromError(err)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Code:    fiber.StatusInternalServerError,
				Message: "服务器内部错误",
				Data:    nil,
			})
		}

		switch sta.Code() {
		case codes.AlreadyExists, codes.InvalidArgument:
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Code:    fiber.StatusBadRequest,
				Message: sta.Message(),
				Data:    nil,
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Code:    fiber.StatusInternalServerError,
				Message: "服务器内部错误",
				Data:    nil,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(view.Response{
		Code:    fiber.StatusOK,
		Message: "注册成功",
		Data: view.RegisterResponse{
			Id:       user.Id,
			Username: user.Username,
		},
	})
}

func (authHandler *AuthHandler) Refresh(ctx *fiber.Ctx) error {
	var req view.RefreshRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    fiber.StatusBadRequest,
			Message: "无效的请求体格式",
			Data:    nil,
		})
	}

	if req.RefreshToken == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    fiber.StatusBadRequest,
			Message: "刷新令牌不能为空",
			Data:    nil,
		})
	}

	c, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	tokenResp, err := authHandler.authClientPool.Get().Refresh(c, &auth.RefreshRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		sta, ok := status.FromError(err)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Code:    fiber.StatusInternalServerError,
				Message: "服务器内部错误",
				Data:    nil,
			})
		}

		switch sta.Code() {
		case codes.Unauthenticated:
			return ctx.Status(fiber.StatusUnauthorized).JSON(view.Response{
				Code:    fiber.StatusUnauthorized,
				Message: sta.Message(),
				Data:    nil,
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Code:    fiber.StatusInternalServerError,
				Message: "服务器内部错误",
				Data:    nil,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(view.Response{
		Code:    fiber.StatusOK,
		Message: "刷新令牌成功",
		Data: view.TokenResponse{
			AccessToken:  tokenResp.AccessToken,
			RefreshToken: tokenResp.RefreshToken,
		},
	})
}

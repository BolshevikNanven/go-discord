package handler

import (
	"context"
	"discord/api/im"
	"discord/app/gateway/internal/client"
	"discord/app/gateway/internal/view"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ImHandler struct {
	imClientPool *client.ImClientPool
}

func NewImHandler(imClientPool *client.ImClientPool) *ImHandler {
	return &ImHandler{
		imClientPool: imClientPool,
	}
}

func (h *ImHandler) SendMessage(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelIdStr := ctx.Params("channel_id")

	var (
		channelId int64
		err       error
	)
	if channelIdStr != "" {
		channelId, err = strconv.ParseInt(channelIdStr, 10, 64)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Code:    -1,
				Message: "Invalid channel id",
			})
		}
	}

	var req view.SendMessageRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Invalid request",
		})
	}

	resp, err := h.imClientPool.Get().SendMessage(context.Background(), &im.SendMessageRequest{
		SpaceId:   req.SpaceId,
		ChannelId: channelId,
		From:      userId,
		To:        req.To,
		Type:      req.Type,
		Content:   req.Content,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "Message sent successfully",
		Data:    resp,
	})
}

func (h *ImHandler) AckMessage(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)

	var req view.AckMessageRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Invalid request",
		})
	}

	resp, err := h.imClientPool.Get().AckMessages(context.Background(), &im.AckMessagesRequest{
		SpaceId:    req.SpaceId,
		UserId:     userId,
		MessageIds: req.MessageIds,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "Message acked successfully",
		Data:    resp,
	})
}

func (h *ImHandler) GetHistory(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelIdStr := ctx.Params("channel_id")

	var req view.PullHistoryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Invalid request",
		})
	}
	if channelIdStr != "" {
		channelId, err := strconv.ParseInt(channelIdStr, 10, 64)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Code:    -1,
				Message: "Invalid channel id",
			})
		}
		req.ChannelId = channelId
	}

	resp, err := h.imClientPool.Get().PullHistory(context.Background(), &im.PullHistoryRequest{
		SpaceId:   req.SpaceId,
		ChannelId: req.ChannelId,
		UserId:    userId,
		From:      req.From,
		Cursor:    req.Cursor,
		Limit:     req.Limit,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "History fetched successfully",
		Data:    resp,
	})
}

func (h *ImHandler) GetInboxMessages(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)

	spaceId := (int64)(ctx.QueryInt("space_id"))
	if spaceId == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}
	limit := (int32)(ctx.QueryInt("limit"))
	if limit == 0 {
		limit = 10
	}

	resp, err := h.imClientPool.Get().GetInboxMessages(context.Background(), &im.GetInboxMessagesRequest{
		SpaceId: spaceId,
		UserId:  userId,
		Limit:   limit,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "Inbox messages fetched successfully",
		Data:    resp,
	})
}

func (h *ImHandler) GetChannelInbox(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelIdStr := ctx.Params("channel_id")
	if channelIdStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Channel id is required",
		})
	}
	channelId, err := strconv.ParseInt(channelIdStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Invalid channel id",
		})
	}

	resp, err := h.imClientPool.Get().GetChannelInbox(context.Background(), &im.GetChannelInboxRequest{
		ChannelId: channelId,
		UserId:    userId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "Channel inbox fetched successfully",
		Data:    resp,
	})
}

func (h *ImHandler) AckChannelMessage(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelIdStr := ctx.Params("channel_id")
	if channelIdStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Channel id is required",
		})
	}
	channelId, err := strconv.ParseInt(channelIdStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Invalid channel id",
		})
	}

	var req view.AckChannelMessageRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "Invalid request",
		})
	}

	resp, err := h.imClientPool.Get().AckChannelMessage(context.Background(), &im.AckChannelMessageRequest{
		ChannelId: channelId,
		UserId:    userId,
		MessageId: req.MessageId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "Channel message acked successfully",
		Data:    resp,
	})
}

package handler

import (
	"discord/api/biz"
	"discord/app/gateway/internal/client"
	"discord/app/gateway/internal/view"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BizHandler struct {
	bizClientPool *client.BizClientPool
}

func NewBizHandler(bizClientPool *client.BizClientPool) *BizHandler {
	return &BizHandler{
		bizClientPool: bizClientPool,
	}
}

func (h *BizHandler) CreateSpace(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)

	request := view.SpaceRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().CreateSpace(ctx.Context(), &biz.CreateSpaceRequest{
		UserId: userId,
		Name:   request.Name,
		Avatar: request.Avatar,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: "创建空间失败",
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "创建成功",
		Data:    response,
	})

}

func (h *BizHandler) JoinSpace(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	spaceId, err := strconv.ParseInt(ctx.Params("space_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().JoinSpace(ctx.Context(), &biz.JoinSpaceRequest{
		UserId:  userId,
		SpaceId: spaceId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "加入成功",
		Data:    response,
	})
}

func (h *BizHandler) LeaveSpace(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	spaceId, err := strconv.ParseInt(ctx.Params("space_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().LeaveSpace(ctx.Context(), &biz.JoinSpaceRequest{
		UserId:  userId,
		SpaceId: spaceId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "离开成功",
		Data:    response,
	})
}

func (h *BizHandler) GetSpaces(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)

	response, err := h.bizClientPool.Get().GetSpaces(ctx.Context(), &biz.GetSpacesRequest{
		UserId: userId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "获取成功",
		Data:    response,
	})
}

func (h *BizHandler) UpdateSpace(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	spaceId, err := strconv.ParseInt(ctx.Params("space_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	request := view.SpaceRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().UpdateSpace(ctx.Context(), &biz.UpdateSpaceRequest{
		UserId:  userId,
		SpaceId: spaceId,
		Name:    request.Name,
		Avatar:  request.Avatar,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "更新成功",
		Data:    response,
	})
}

func (h *BizHandler) DeleteSpace(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	spaceId, err := strconv.ParseInt(ctx.Params("space_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().DeleteSpace(ctx.Context(), &biz.DeleteSpaceRequest{
		UserId:  userId,
		SpaceId: spaceId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "删除成功",
		Data:    response,
	})
}

func (h *BizHandler) CreateChannel(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)

	request := view.CreateChannelRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().CreateChannel(ctx.Context(), &biz.CreateChannelRequest{
		UserId:  userId,
		SpaceId: request.SpaceID,
		Name:    request.Name,
		Type:    request.Type,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "创建成功",
		Data:    response,
	})
}

func (h *BizHandler) JoinChannel(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelId, err := strconv.ParseInt(ctx.Params("channel_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().JoinChannel(ctx.Context(), &biz.JoinChannelRequest{
		UserId:    userId,
		ChannelId: channelId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "加入成功",
		Data:    response,
	})
}

func (h *BizHandler) LeaveChannel(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelId, err := strconv.ParseInt(ctx.Params("channel_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().LeaveChannel(ctx.Context(), &biz.JoinChannelRequest{
		UserId:    userId,
		ChannelId: channelId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "离开成功",
		Data:    response,
	})
}

func (h *BizHandler) GetChannels(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	spaceId := (int64)(ctx.QueryInt("space_id"))

	if spaceId == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().GetChannels(ctx.Context(), &biz.GetChannelsRequest{
		UserId:  userId,
		SpaceId: spaceId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "获取成功",
		Data:    response,
	})
}

func (h *BizHandler) UpdateChannel(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelId, err := strconv.ParseInt(ctx.Params("channel_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	request := view.UpdateChannelRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().UpdateChannel(ctx.Context(), &biz.UpdateChannelRequest{
		UserId:    userId,
		ChannelId: channelId,
		Name:      request.Name,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "更新成功",
		Data:    response,
	})
}

func (h *BizHandler) DeleteChannel(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int64)
	channelId, err := strconv.ParseInt(ctx.Params("channel_id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
			Code:    -1,
			Message: "请求参数错误",
		})
	}

	response, err := h.bizClientPool.Get().DeleteChannel(ctx.Context(), &biz.DeleteChannelRequest{
		UserId:    userId,
		ChannelId: channelId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
			Code:    -1,
			Message: err.Error(),
		})
	}

	return ctx.JSON(view.Response{
		Code:    0,
		Message: "删除成功",
		Data:    response,
	})
}

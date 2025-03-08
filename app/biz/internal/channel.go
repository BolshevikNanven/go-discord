package internal

import (
	"context"
	"discord/api/biz"
	"discord/app/biz/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetChannels(ctx context.Context, req *biz.GetChannelsRequest) (*biz.ChannelList, error) {
	isMember, err := s.spaceUserRepo.IsSpaceMember(req.SpaceId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !isMember {
		return nil, status.Error(codes.PermissionDenied, "无权限")
	}

	channels, err := s.channelRepo.List(req.SpaceId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	bizChannels := make([]*biz.Channel, len(channels))
	for i, channel := range channels {
		bizChannels[i] = &biz.Channel{
			Id:    channel.Id,
			Name:  channel.Name,
			Owner: channel.Owner,
			Type:  channel.Type,
		}
	}

	return &biz.ChannelList{
		Channels: bizChannels,
	}, nil

}
func (s *Server) CreateChannel(ctx context.Context, req *biz.CreateChannelRequest) (*biz.Channel, error) {
	var err error

	isMember, err := s.spaceUserRepo.IsSpaceMember(req.SpaceId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !isMember {
		return nil, status.Error(codes.PermissionDenied, "无权限")
	}

	channel := model.Channel{
		SpaceId: req.SpaceId,
		Name:    req.Name,
		Owner:   req.UserId,
	}

	switch req.Type {
	case "PUBLIC":
		channel.Type = "PUBLIC"
	default:
		channel.Type = "PRIVATE"
	}

	channel, err = s.channelRepo.Create(channel)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.Channel{
		Id:    channel.Id,
		Name:  channel.Name,
		Owner: channel.Owner,
		Type:  channel.Type,
	}, nil
}

func (s *Server) UpdateChannel(ctx context.Context, req *biz.UpdateChannelRequest) (*biz.Channel, error) {
	preChannel, err := s.channelRepo.GetByID(req.ChannelId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if preChannel.Owner != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "无权限")
	}

	channel := model.Channel{
		Name: req.Name,
	}
	channel, err = s.channelRepo.Update(channel)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &biz.Channel{
		Id:   channel.Id,
		Name: channel.Name,
	}, nil
}

func (s *Server) DeleteChannel(ctx context.Context, req *biz.DeleteChannelRequest) (*biz.SuccessResponse, error) {
	preChannel, err := s.channelRepo.GetByID(req.ChannelId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if preChannel.Owner != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "无权限")
	}

	err = s.channelRepo.Delete(req.ChannelId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.SuccessResponse{
		Success: true,
	}, nil

}

func (s *Server) JoinChannel(ctx context.Context, req *biz.JoinChannelRequest) (*biz.Channel, error) {
	isMember, err := s.channelUserRepo.IsChannelMember(req.ChannelId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if isMember {
		return nil, status.Error(codes.PermissionDenied, "已加入频道")
	}

	if err := s.channelUserRepo.Create(req.ChannelId, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	channel, _ := s.channelRepo.GetByID(req.ChannelId)

	return &biz.Channel{
		Id:    channel.Id,
		Name:  channel.Name,
		Owner: channel.Owner,
		Type:  channel.Type,
	}, nil
}

func (s *Server) LeaveChannel(ctx context.Context, req *biz.JoinChannelRequest) (*biz.SuccessResponse, error) {
	isMember, err := s.channelUserRepo.IsChannelMember(req.ChannelId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !isMember {
		return nil, status.Error(codes.PermissionDenied, "未加入频道")
	}

	err = s.channelUserRepo.Delete(req.ChannelId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.SuccessResponse{
		Success: true,
	}, nil
}

func (s *Server) GetChannelIds(ctx context.Context, req *biz.GetChannelIdsRequest) (*biz.GetChannelIdsResponse, error) {
	isMember, err := s.spaceUserRepo.IsSpaceMember(req.SpaceId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !isMember {
		return nil, status.Error(codes.PermissionDenied, "无权限")
	}

	channelIds, err := s.channelRepo.All(req.SpaceId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.GetChannelIdsResponse{
		ChannelIds: channelIds,
	}, nil
}

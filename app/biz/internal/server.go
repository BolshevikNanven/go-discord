package internal

import (
	"discord/api/biz"
	"discord/app/biz/internal/repository"
)

type Server struct {
	biz.UnimplementedBizServiceServer
	spaceRepo       repository.SpaceRepository
	channelRepo     repository.ChannelRepository
	spaceUserRepo   repository.SpaceUserRepository
	channelUserRepo repository.ChannelUserRepository
}

func NewServer(
	spaceRepository repository.SpaceRepository,
	channelRepository repository.ChannelRepository,
	spaceUserRepository repository.SpaceUserRepository,
	channelUserRepository repository.ChannelUserRepository,
) biz.BizServiceServer {
	return &Server{
		spaceRepo:       spaceRepository,
		channelRepo:     channelRepository,
		spaceUserRepo:   spaceUserRepository,
		channelUserRepo: channelUserRepository,
	}
}

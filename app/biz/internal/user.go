package internal

import (
	"context"
	"discord/api/biz"
)

func (s *Server) IsSpaceMember(ctx context.Context, req *biz.IsSpaceMemberRequest) (*biz.IsSpaceMemberResponse, error) {
	ok, err := s.spaceUserRepo.IsSpaceMember(req.SpaceId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &biz.IsSpaceMemberResponse{
		IsMember: ok,
	}, nil
}

func (s *Server) IsChannelMember(ctx context.Context, req *biz.IsChannelMemberRequest) (*biz.IsChannelMemberResponse, error) {
	ok, err := s.channelUserRepo.IsChannelMember(req.ChannelId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &biz.IsChannelMemberResponse{
		IsMember: ok,
	}, nil
}

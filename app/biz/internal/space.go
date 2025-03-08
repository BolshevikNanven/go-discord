package internal

import (
	"context"
	"discord/api/biz"
	"discord/app/biz/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSpaces(ctx context.Context, req *biz.GetSpacesRequest) (*biz.SpaceList, error) {
	spaces, err := s.spaceRepo.UserList(req.UserId)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	bizSpaces := make([]*biz.Space, len(spaces))
	for i, space := range spaces {
		bizSpaces[i] = &biz.Space{
			Id:     space.Id,
			Name:   space.Name,
			Avatar: space.Avatar,
			Owner:  space.Owner,
		}
	}

	return &biz.SpaceList{
		Spaces: bizSpaces,
	}, nil

}

func (s *Server) CreateSpace(ctx context.Context, req *biz.CreateSpaceRequest) (*biz.Space, error) {
	space := model.Space{
		Name:   req.Name,
		Avatar: req.Avatar,
		Owner:  req.UserId,
	}

	space, err := s.spaceRepo.Create(space)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.Space{
		Id:     space.Id,
		Name:   space.Name,
		Avatar: space.Avatar,
		Owner:  space.Owner,
	}, nil
}

func (s *Server) UpdateSpace(ctx context.Context, req *biz.UpdateSpaceRequest) (*biz.Space, error) {
	// 检查space权限
	pre, err := s.spaceRepo.GetByID(req.SpaceId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "space不存在")
	}
	if pre.Owner != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "无权限")
	}

	space := model.Space{
		Id:     req.SpaceId,
		Name:   req.Name,
		Avatar: req.Avatar,
	}

	if err := s.spaceRepo.Update(space); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.Space{
		Id:     space.Id,
		Name:   space.Name,
		Avatar: space.Avatar,
		Owner:  space.Owner,
	}, nil
}

func (s *Server) DeleteSpace(ctx context.Context, req *biz.DeleteSpaceRequest) (*biz.SuccessResponse, error) {
	// 检查space权限
	pre, err := s.spaceRepo.GetByID(req.SpaceId)
	if err != nil {

		return nil, status.Error(codes.NotFound, "space不存在")
	}
	if pre.Owner != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "无权限")
	}

	if err := s.spaceRepo.Delete(req.SpaceId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.SuccessResponse{
		Success: true,
	}, nil
}

func (s *Server) JoinSpace(ctx context.Context, req *biz.JoinSpaceRequest) (*biz.Space, error) {
	isMember, err := s.spaceUserRepo.IsSpaceMember(req.SpaceId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if isMember {
		return nil, status.Error(codes.PermissionDenied, "已加入空间")
	}

	if err := s.spaceUserRepo.Create(req.SpaceId, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	space, _ := s.spaceRepo.GetByID(req.SpaceId)

	return &biz.Space{
		Id:     req.SpaceId,
		Name:   space.Name,
		Avatar: space.Avatar,
		Owner:  space.Owner,
	}, nil

}
func (s *Server) LeaveSpace(ctx context.Context, req *biz.JoinSpaceRequest) (*biz.SuccessResponse, error) {
	isMember, err := s.spaceUserRepo.IsSpaceMember(req.SpaceId, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !isMember {
		return nil, status.Error(codes.PermissionDenied, "未加入空间")
	}

	if err := s.spaceUserRepo.Delete(req.SpaceId, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.SuccessResponse{
		Success: true,
	}, nil
}

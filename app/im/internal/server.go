package internal

import (
	"context"
	"discord/api/biz"
	"discord/api/im"
	"discord/app/im/internal/client"
	"discord/app/im/internal/model"
	"discord/app/im/internal/repository"
	"discord/pkg/snowflakeutil"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	im.UnimplementedImServiceServer
	messageRepo repository.MessageRepository
	inboxRepo   repository.InboxRepository
	mqRepo      repository.MqRepository

	bizClientPool *client.BizClientPool

	idGenerator *snowflakeutil.Node
}

func NewServer(
	messageRepo repository.MessageRepository,
	inboxRepo repository.InboxRepository,
	mqRepo repository.MqRepository,
	bizClientPool *client.BizClientPool,
	idGenerator *snowflakeutil.Node,
) im.ImServiceServer {
	return &Server{
		messageRepo:   messageRepo,
		inboxRepo:     inboxRepo,
		mqRepo:        mqRepo,
		bizClientPool: bizClientPool,
		idGenerator:   idGenerator,
	}
}

func (s *Server) SendMessage(ctx context.Context, req *im.SendMessageRequest) (*im.SendMessageResponse, error) {
	resp, err := s.bizClientPool.Get().IsSpaceMember(ctx, &biz.IsSpaceMemberRequest{
		SpaceId: req.SpaceId,
		UserId:  req.From,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check space member")
	}
	if !resp.IsMember {
		return nil, status.Errorf(codes.PermissionDenied, "not a member of the space")
	}

	message := model.ChatMessage{
		Id:        s.idGenerator.GenerateID(),
		SpaceId:   req.SpaceId,
		From:      req.From,
		To:        req.To,
		Content:   req.Content,
		Type:      req.Type,
		CreatedAt: time.Now(),
	}
	isChannel := (req.ChannelId != 0)
	if isChannel {
		resp, err := s.bizClientPool.Get().IsChannelMember(ctx, &biz.IsChannelMemberRequest{
			ChannelId: req.ChannelId,
			UserId:    req.From,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to check channel member")
		}
		if !resp.IsMember {
			return nil, status.Errorf(codes.PermissionDenied, "not a member of the channel")
		}
		message.To = req.ChannelId
	}

	// 消息入队
	if err := s.mqRepo.SendChatMessage(&message, isChannel); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send message to mq")
	}

	return &im.SendMessageResponse{
		MessageId: message.Id,
	}, nil
}

func (s *Server) AckMessages(ctx context.Context, req *im.AckMessagesRequest) (*im.AckMessagesResponse, error) {
	resp, err := s.bizClientPool.Get().IsSpaceMember(ctx, &biz.IsSpaceMemberRequest{
		SpaceId: req.SpaceId,
		UserId:  req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check space member")
	}
	if !resp.IsMember {
		return nil, status.Errorf(codes.PermissionDenied, "not a member of the space")
	}

	err = s.inboxRepo.RemoveInboxMessages(req.UserId, req.SpaceId, req.MessageIds)
	if err != nil {
		fmt.Println(err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove inbox message")
	}

	return &im.AckMessagesResponse{
		Success: true,
	}, nil
}

func (s *Server) PullHistory(ctx context.Context, req *im.PullHistoryRequest) (*im.PullHistoryResponse, error) {
	resp, err := s.bizClientPool.Get().IsSpaceMember(ctx, &biz.IsSpaceMemberRequest{
		SpaceId: req.SpaceId,
		UserId:  req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check space member")
	}
	if !resp.IsMember {
		return nil, status.Errorf(codes.PermissionDenied, "not a member of the space")
	}

	// 群聊
	if req.ChannelId != 0 {
		resp, err := s.bizClientPool.Get().IsChannelMember(ctx, &biz.IsChannelMemberRequest{
			ChannelId: req.ChannelId,
			UserId:    req.UserId,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to check channel member")
		}
		if !resp.IsMember {
			return nil, status.Errorf(codes.PermissionDenied, "not a member of the channel")
		}

		messages, err := s.messageRepo.GetChannelMessages(req.ChannelId, req.SpaceId, req.Cursor, int(req.Limit))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get channel messages")
		}

		var pbMessages []*im.Message
		for _, message := range messages {
			pbMessages = append(pbMessages, &im.Message{
				SpaceId:   message.SpaceId,
				Id:        message.Id,
				Content:   message.Content,
				From:      message.From,
				To:        message.To,
				Type:      message.Type,
				CreatedAt: message.CreatedAt.Unix(),
			})
		}
		return &im.PullHistoryResponse{
			Messages: pbMessages,
		}, nil
	}

	// 单聊
	messages, err := s.messageRepo.GetMessages(req.UserId, req.SpaceId, req.From, req.Cursor, int(req.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get messages")
	}

	var pbMessages []*im.Message
	for _, message := range messages {
		pbMessages = append(pbMessages, &im.Message{
			SpaceId:   message.SpaceId,
			Id:        message.Id,
			Content:   message.Content,
			From:      message.From,
			To:        message.To,
			Type:      message.Type,
			CreatedAt: message.CreatedAt.Unix(),
		})
	}

	return &im.PullHistoryResponse{
		Messages: pbMessages,
	}, nil
}

func (s *Server) GetInboxMessages(ctx context.Context, req *im.GetInboxMessagesRequest) (*im.GetInboxMessagesResponse, error) {
	resp, err := s.bizClientPool.Get().IsSpaceMember(ctx, &biz.IsSpaceMemberRequest{
		SpaceId: req.SpaceId,
		UserId:  req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check space member")
	}
	if !resp.IsMember {
		return nil, status.Errorf(codes.PermissionDenied, "not a member of the space")
	}

	messages, err := s.inboxRepo.GetInboxMessages(req.UserId, req.SpaceId, int(req.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbMessages []*im.Message
	for _, message := range messages {
		pbMessages = append(pbMessages, &im.Message{
			SpaceId:   message.SpaceId,
			Id:        message.Id,
			Content:   message.Content,
			From:      message.From,
			To:        message.To,
			Type:      message.Type,
			CreatedAt: message.CreatedAt.Unix(),
		})
	}

	return &im.GetInboxMessagesResponse{
		Messages: pbMessages,
	}, nil
}

func (s *Server) AckChannelMessage(ctx context.Context, req *im.AckChannelMessageRequest) (*im.AckChannelMessageResponse, error) {
	resp, err := s.bizClientPool.Get().IsChannelMember(ctx, &biz.IsChannelMemberRequest{
		ChannelId: req.ChannelId,
		UserId:    req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check channel member")
	}
	if !resp.IsMember {
		return nil, status.Errorf(codes.PermissionDenied, "not a member of the channel")
	}

	err = s.inboxRepo.UpdateChannelAckMessageId(req.ChannelId, req.UserId, req.MessageId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update channel ack message id")
	}

	return &im.AckChannelMessageResponse{
		Success: true,
	}, nil

}

func (s *Server) GetChannelInbox(ctx context.Context, req *im.GetChannelInboxRequest) (*im.GetChannelInboxResponse, error) {
	resp, err := s.bizClientPool.Get().IsChannelMember(ctx, &biz.IsChannelMemberRequest{
		ChannelId: req.ChannelId,
		UserId:    req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check channel member")
	}
	if !resp.IsMember {
		return nil, status.Errorf(codes.PermissionDenied, "not a member of the channel")
	}

	current, last, err := s.inboxRepo.GetChannelInbox(req.ChannelId, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &im.GetChannelInboxResponse{
		Current: current,
		Last:    last,
	}, nil
}

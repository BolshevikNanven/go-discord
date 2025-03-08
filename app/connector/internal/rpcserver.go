package internal

import (
	"context"
	"discord/api/connector"
	"discord/app/connector/internal/hub"
)

type RpcServer struct {
	connector.UnimplementedConnectorServiceServer
	hub *hub.Hub
}

func NewRpcServer(hub *hub.Hub) connector.ConnectorServiceServer {
	return &RpcServer{
		hub: hub,
	}
}

func (s *RpcServer) SendMessage(ctx context.Context, req *connector.SendMessageRequest) (*connector.SendMessageResponse, error) {
	success := s.hub.SendMessage(req.UserId, []byte(req.Message))
	return &connector.SendMessageResponse{
		Success: success,
	}, nil
}

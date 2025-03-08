package internal

import (
	"context"
	"discord/api/auth"
	"discord/app/auth/internal/repository"
	"discord/pkg/jwtutil"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	auth.UnimplementedAuthServiceServer
	userRepo  repository.UserRepository
	jwtConfig *jwtutil.Config
	logger    *zap.Logger
}

func NewServer(userRepo repository.UserRepository, jwtConfig *jwtutil.Config, logger *zap.Logger) auth.AuthServiceServer {
	return &Server{
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
		logger:    logger,
	}
}

func (s *Server) Login(ctx context.Context, req *auth.LoginRequest) (*auth.TokenResponse, error) {
	s.logger.Info("login", zap.String("username", req.Username), zap.String("password", req.Password))
	user := s.userRepo.GetByUsername(req.Username)
	if user == nil {
		return nil, status.Error(codes.NotFound, "用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.InvalidArgument, "用户名或密码错误")
	}

	accessToken, err := jwtutil.GenerateToken(int64(user.Id), jwtutil.AccessToken, s.jwtConfig)
	if err != nil {
		return nil, status.Error(codes.Internal, "生成访问令牌失败")
	}

	refreshToken, err := jwtutil.GenerateToken(int64(user.Id), jwtutil.RefreshToken, s.jwtConfig)
	if err != nil {
		return nil, status.Error(codes.Internal, "生成刷新令牌失败")
	}

	return &auth.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.TokenResponse, error) {
	// 验证刷新令牌
	claims, err := jwtutil.ValidateToken(req.RefreshToken, jwtutil.RefreshToken, s.jwtConfig)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "无效的刷新令牌")
	}

	// 生成新的访问令牌和刷新令牌
	accessToken, err := jwtutil.GenerateToken(claims.UserId, jwtutil.AccessToken, s.jwtConfig)
	if err != nil {
		return nil, status.Error(codes.Internal, "生成访问令牌失败")
	}

	return &auth.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: req.RefreshToken,
	}, nil
}

func (s *Server) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.User, error) {
	// 检查用户名是否已存在
	existingUser := s.userRepo.GetByUsername(req.Username)
	if existingUser != nil {
		return nil, status.Error(codes.AlreadyExists, "用户名已存在")
	}

	// 对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "密码加密失败")
	}

	// 创建新用户
	user, err := s.userRepo.Create(req.Username, string(hashedPassword))
	if err != nil {
		return nil, status.Error(codes.Internal, "创建用户失败")
	}

	return &auth.User{
		Id:       int64(user.Id),
		Username: user.Username,
	}, nil
}

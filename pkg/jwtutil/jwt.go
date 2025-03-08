package jwtutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type Claims struct {
	UserId int64     `json:"user_id"`
	Type   TokenType `json:"type"`
	jwt.RegisteredClaims
}

type Config struct {
	SecretKey          string        `yaml:"secretKey"`
	AccessTokenExpiry  time.Duration `yaml:"accessTokenExpiry"`
	RefreshTokenExpiry time.Duration `yaml:"refreshTokenExpiry"`
}

// GenerateToken 生成JWT令牌
func GenerateToken(userId int64, tokenType TokenType, config *Config) (string, error) {
	var expiry time.Duration
	switch tokenType {
	case AccessToken:
		expiry = config.AccessTokenExpiry
	case RefreshToken:
		expiry = config.RefreshTokenExpiry
	default:
		return "", fmt.Errorf("无效的令牌类型")
	}

	claims := Claims{
		UserId: userId,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string, config *Config) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return nil, fmt.Errorf("解析令牌失败: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的令牌")
}

// ValidateToken 验证令牌的有效性
func ValidateToken(tokenString string, expectedType TokenType, config *Config) (*Claims, error) {
	claims, err := ParseToken(tokenString, config)
	if err != nil {
		return nil, err
	}

	if claims.Type != expectedType {
		return nil, fmt.Errorf("令牌类型不匹配")
	}

	return claims, nil
}

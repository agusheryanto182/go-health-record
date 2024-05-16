package jwt

import (
	"time"

	"github.com/agusheryanto182/go-health-record/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTInterface interface {
	GenerateJWT(ID, phone_number string) (string, error)
	ValidateToken(tokenString string) (*JWTPayload, error)
}

type JWTService struct {
	cfg *config.Global
}

func NewJWTService(cfg *config.Global) JWTInterface {
	return &JWTService{
		cfg: cfg,
	}
}

type JWTPayload struct {
	Id   string
	Role string
}

type JWTClaims struct {
	Id   string
	Role string
	jwt.RegisteredClaims
}

func (s *JWTService) GenerateJWT(id, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
	})

	tokenString, err := token.SignedString([]byte(s.cfg.Jwt.Secret))
	return tokenString, err
}

func (s *JWTService) ValidateToken(tokenString string) (*JWTPayload, error) {
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
		return nil, err
	}

	payload := &JWTPayload{
		Id:   claims.Id,
		Role: claims.Role,
	}

	return payload, nil
}

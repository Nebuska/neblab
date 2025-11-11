package jwtAuth

import (
	"github.com/Nebuska/task-tracker/config"
	"time"
)
import "github.com/golang-jwt/jwt/v5"

type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

type UserClaims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

type JWTToken string

func NewJWTManager(config *config.Config) *JWTManager {
	return &JWTManager{
		secretKey:     []byte(config.JWTSecret),
		tokenDuration: config.JWTExpire,
	}
}

func (j *JWTManager) Generate(userId uint) (JWTToken, error) {
	claims := UserClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "TaskTracker",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.secretKey)

	return JWTToken(tokenString), err
}

func (j *JWTManager) Verify(rawToken JWTToken) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		string(rawToken),
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

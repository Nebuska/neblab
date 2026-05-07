package session

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserId    uint
	Token     Token
	PrevToken Token
	ExpiresAt time.Time
	UserAgent string
	IpAddress string
}

func NewSession(userId uint, userAgent, IpAddress string) Session {
	token := GenerateToken()
	return Session{
		UserId:    userId,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour),
		UserAgent: userAgent,
		IpAddress: IpAddress,
	}
}

func (s Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

type Token string

// GenerateToken creates a new token
func GenerateToken() Token {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return Token(base64.RawURLEncoding.EncodeToString(b))
}

func Tokenize(userToken string) Token {
	return Token(userToken)
}

func (t Token) GetRedisKey() string {
	return "session:" + string(t)
}

func (t Token) String() string {
	return string(t)
}

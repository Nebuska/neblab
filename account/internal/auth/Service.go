package auth

import (
	"context"
	"errors"

	"github.com/Nebuska/neblab/account/internal/auth/dto"
	"github.com/Nebuska/neblab/account/internal/credentials"
	"github.com/Nebuska/neblab/account/internal/session"
	"github.com/Nebuska/neblab/account/internal/user"
	"github.com/Nebuska/neblab/shared/jwtAuth"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, data dto.RegisterData) error

	Logout(ctx context.Context, session session.Session) error
	Login(ctx context.Context, data dto.LoginData) (session.Session, error)
	Refresh(ctx context.Context, session session.Session) (jwtAuth.JWTToken, error)
	OneTime(ctx context.Context, data dto.LoginData) (jwtAuth.JWTToken, error)
}

type service struct {
	credRepo credentials.Repository
	userRepo user.Repository
	jwt      *jwtAuth.JWTManager
	db       *gorm.DB
}

func NewAuthService(
	credRepo credentials.Repository,
	userRepo user.Repository,
	jwt *jwtAuth.JWTManager,
	db *gorm.DB,
) Service {
	return &service{
		credRepo: credRepo,
		userRepo: userRepo,
		jwt:      jwt,
		db:       db,
	}
}

func (s *service) Register(ctx context.Context, data dto.RegisterData) error {
	_, isExist, err := s.credRepo.CheckEmail(ctx, nil, data.Email)
	if err != nil {
		return err
	}
	if isExist {
		return dto.ErrEmailAlreadyExists
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	cred := credentials.Credentials{
		Email:    data.Email,
		Username: data.Username,
		Password: string(passHash),
	}
	usr := user.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}
	err = s.db.Transaction(func(tx *gorm.DB) error {
		usr, err = s.userRepo.Create(ctx, tx, usr)
		if err != nil {
			return err
		}
		cred.ID = usr.ID
		cred, err = s.credRepo.Create(ctx, tx, cred)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *service) Login(ctx context.Context, data dto.LoginData) (session.Session, error) {
	//todo : Sessions not implemented
	return session.Session{}, errors.New("not implemented")
}

func (s *service) Refresh(ctx context.Context, session session.Session) (jwtAuth.JWTToken, error) {
	//todo : Sessions not implemented
	return "", errors.New("not implemented")
}

func (s *service) Logout(ctx context.Context, session session.Session) error {
	//todo : Sessions not implemented
	return errors.New("not implemented")
}

func (s *service) OneTime(ctx context.Context, data dto.LoginData) (jwtAuth.JWTToken, error) {
	cred, err := s.credRepo.Find(ctx, nil, data.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(cred.Password), []byte(data.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.Generate(cred.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

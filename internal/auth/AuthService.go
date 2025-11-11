package auth

import (
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
	"github.com/Nebuska/task-tracker/pkg/jwtAuth"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(username, email, password string) (jwtAuth.JWTToken, error)
	Login(username, password string) (jwtAuth.JWTToken, error)
}

type authService struct {
	repo Repository
	jwt  *jwtAuth.JWTManager
}

func NewAuthService(repo Repository, jwt *jwtAuth.JWTManager) Service {
	return &authService{repo: repo, jwt: jwt}
}

func (service *authService) Register(username, email, password string) (jwtAuth.JWTToken, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", appError.New(errorCodes.InternalError, "bcrypt", err.Error())
	}
	credentials, err := service.repo.Register(username, email, string(passHash))
	if err != nil {
		return "", err
	}
	token, err := service.jwt.Generate(credentials.ID)
	if err != nil {
		return "", appError.New(errorCodes.InternalError, "jwt", err.Error())
	}

	return token, nil
}

func (service *authService) Login(username, password string) (jwtAuth.JWTToken, error) {
	userCredentials, err := service.repo.GetCredentials(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userCredentials.Password), []byte(password))
	if err != nil {
		return "", appError.New(errorCodes.InternalError, "bcrypt", err.Error())
	}
	token, err := service.jwt.Generate(userCredentials.ID)
	if err != nil {
		return "", appError.New(errorCodes.InternalError, "jwt", err.Error())
	}
	return token, nil
}

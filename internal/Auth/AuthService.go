package Auth

import (
	"task-tracker/pkg/jwtAuth"

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

func (service *authService) Register(username, email, password string) (token jwtAuth.JWTToken, err error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	credentials, err := service.repo.Register(username, email, string(passHash))
	if err != nil {
		return
	}

	return service.jwt.Generate(credentials.ID)
}

func (service *authService) Login(username, password string) (token jwtAuth.JWTToken, err error) {
	userCredentials, err := service.repo.GetCredentials(username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userCredentials.Password), []byte(password))
	if err != nil {
		return
	}

	return service.jwt.Generate(userCredentials.ID)
}

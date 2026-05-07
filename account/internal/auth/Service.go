package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Nebuska/neblab/account/internal/credentials"
	"github.com/Nebuska/neblab/account/internal/dto"
	"github.com/Nebuska/neblab/account/internal/session"
	"github.com/Nebuska/neblab/account/internal/user"
	"github.com/Nebuska/neblab/shared/jwtAuth"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

// Service interface is for Auth service that can be accessed by outside
type Service interface {
	// Register takes dto.RegisterData and parses it to create credentials.Credentials and user.User
	//
	// takes unencrypted password and encrypts it than register new user in database
	//
	// returns dto.ErrUserAlreadyExist if username used before
	// returns bcrypt.ErrPasswordTooLong if password is longer than 72 byte
	// returns gorm error or bcrypt error on internal error
	Register(ctx context.Context, data dto.RegisterData) error

	// Login takes dto.LoginData and verify the user than create new session for user
	//
	// returns dto.ErrUserNotFound if no user in given username
	// returns dto.ErrWrongPassword if password is wrong
	// returns gorm error or bcrypt error on internal error
	Login(ctx context.Context, data dto.LoginData) (session.Token, error)
	Refresh(ctx context.Context, token session.Token) (session.Token, error)
	Logout(ctx context.Context, token session.Token) error
	GetJwt(ctx context.Context, token session.Token) (jwtAuth.JWTToken, error)
	OneTime(ctx context.Context, data dto.LoginData) (jwtAuth.JWTToken, error)
}

type service struct {
	credRepo    credentials.Repository
	userRepo    user.Repository
	sessionRepo session.Repository
	jwt         *jwtAuth.JWTManager
	db          *gorm.DB
}

func NewAuthService(
	credRepo credentials.Repository,
	userRepo user.Repository,
	sessionRepo session.Repository,
	jwt *jwtAuth.JWTManager,
	db *gorm.DB,
) Service {
	return &service{
		credRepo:    credRepo,
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwt:         jwt,
		db:          db,
	}
}

// Register takes dto.RegisterData and parses it to create credentials.Credentials and user.User
//
// takes unencrypted password and encrypts it than register new user in database
//
// returns dto.ErrUserAlreadyExist if username used before
// returns bcrypt.ErrPasswordTooLong if password is longer than 72 byte
// returns gorm error or bcrypt error on internal error
func (s *service) Register(ctx context.Context, data dto.RegisterData) error {
	/* todo: prior email check could reduce run time in fails
	_, isExist, err := s.credRepo.CheckEmail(ctx, nil, data.Email)
	if err != nil {
		return err
	}
	*/
	data.Password = norm.NFC.String(data.Password)
	passHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		usr, err := s.userRepo.Create(ctx, tx, user.FromRegisterDTO(data))
		if err != nil {
			return err
		}
		cred := credentials.FromRegisterDTO(data, string(passHash))
		cred.ID = usr.ID
		_, err = s.credRepo.Create(ctx, tx, cred)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// Login takes dto.LoginData and verify the user than create new session for user
//
// returns dto.ErrUserNotFound if no user in given username
// returns dto.ErrWrongPassword if password is wrong
// returns gorm error or bcrypt error on internal error
func (s *service) Login(ctx context.Context, data dto.LoginData) (session.Token, error) {
	cred, err := s.credRepo.Find(ctx, nil, data.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(cred.Password), []byte(data.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", dto.ErrWrongPassword
		}
		return "", err
	}
	// todo: a middleware in api should insert context values
	ipAddress := ctx.Value(dto.CKEY_IpAddress).(string)
	userAgents := ctx.Value(dto.CKEY_UserAgent).(string)
	ses := session.NewSession(cred.ID, ipAddress, userAgents)
	ses, err = s.sessionRepo.Create(ctx, ses)
	for err != nil && errors.Is(err, dto.ErrSessionTokenConflict) {
		ses.Token = session.GenerateToken()
		ses, err = s.sessionRepo.Create(ctx, ses)
	}
	if err != nil {
		return "", err
	}
	return ses.Token, nil
}

// Refresh changes expired session.Token and resumes the session.Session
func (s *service) Refresh(ctx context.Context, token session.Token) (session.Token, error) {
	ses, err := s.sessionRepo.GetByToken(ctx, token)
	if err != nil {
		return "", err
	}
	if condition, _ := s.ValidateSession(ctx, &ses); condition {
		return "", dto.ErrSessionMismatch
	}
	newToken := session.GenerateToken()
	ses, err = s.sessionRepo.RefreshToken(ctx, ses, newToken, 60*time.Minute)
	for err != nil && errors.Is(err, dto.ErrSessionTokenConflict) {
		newToken = session.GenerateToken()
		ses, err = s.sessionRepo.RefreshToken(ctx, ses, newToken, 60*time.Minute)
	}
	if err != nil {
		return "", err
	}
	return ses.Token, nil
}

// Logout ends the session.Session
func (s *service) Logout(ctx context.Context, token session.Token) error {
	err := s.sessionRepo.Delete(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

// GetJwt takes a session.Token and if session isn't expired returns a 5 minute JwtToken
func (s *service) GetJwt(ctx context.Context, token session.Token) (jwtAuth.JWTToken, error) {
	ses, err := s.sessionRepo.GetByToken(ctx, token)
	if err != nil {
		return "", err
	}
	jwtToken, err := s.jwt.Generate(ses.UserId)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

// OneTime creates a JwtToken for limited time outside of sessions
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

// ValidateSession checks stability of session by ip address and user agents
//
// it returns true if session is considered healthy
// it will also return a string explaining reason why it isn't stable even if healthy
//
// Will try to change the Session.IpAddress if they mismatch
func (s *service) ValidateSession(ctx context.Context, ses *session.Session) (bool, string) {
	ipAddress := ctx.Value(dto.CKEY_IpAddress).(string)
	userAgents := ctx.Value(dto.CKEY_UserAgent).(string)
	if ses.IpAddress != ipAddress {
		//Could be network change need to check physical location
		ses.IpAddress = ipAddress
		return true, "IpAddress not match"
	}
	if ses.UserAgent != userAgents {
		return false, "UserAgent not match"
	}
	return true, ""
}

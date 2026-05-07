package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Nebuska/neblab/account/internal/dto"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Repository is Session repository that handles database requests for Services
type Repository interface {

	// Create adds a new entry of Session to redis and database
	//
	// returns redis or gorm error on internal error
	Create(ctx context.Context, session Session) (Session, error)

	// GetByToken takes session.Token and returns the session.Session
	//
	// Checks redis first for session if not found it will check the database
	//
	// returns dto.ErrSessionNotFound if no session in database corresponding to token
	// returns redis or gorm error on internal error
	GetByToken(ctx context.Context, sessionToken Token) (Session, error)

	// RefreshToken changes the access session.Token for a Session and increases lifespan of session
	//
	// Rewrites Session.PrevToken to oldToken for possible security check
	// Rewrites Session.ExpiresAt to Now + duration
	//
	// returns dto.ErrSessionNotFound if oldToken is not a known Session.Token
	// returns gorm or redis error on internal errors
	RefreshToken(ctx context.Context, session Session, newToken Token, duration time.Duration) (Session, error)

	// Delete takes session.Token and closes that session permanently
	//
	// returns redis or gorm error on internal error
	Delete(ctx context.Context, token Token) error
}

type repository struct {
	db  *gorm.DB
	rdb redis.Client
}

func NewRepository(db *gorm.DB, rdb redis.Client) Repository {
	return &repository{
		db:  db,
		rdb: rdb,
	}
}

// WithTx is used for creating transactions
// for making sure when a failure happens all database actions will be cancelled.
//
// While not using transactions tx parameter should be nil
func (r *repository) WithTx(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

// Create adds a new entry of Session to redis and database
//
// returns dto.ErrSessionTokenConflict if session.Token is used
// returns redis or gorm error on internal error
func (r *repository) Create(ctx context.Context, session Session) (Session, error) {
	err := r.db.WithContext(ctx).Create(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return Session{}, dto.ErrSessionTokenConflict
		}
		return Session{}, err
	}
	err = r.setSessionOnRedis(ctx, session)
	return session, err
}

// GetByToken takes session.Token and returns the session.Session
//
// # Checks redis first for session if not found it will check the database
//
// returns dto.ErrSessionNotFound if no session in database corresponding to token
// returns redis or gorm error on internal error
func (r *repository) GetByToken(ctx context.Context, sessionToken Token) (Session, error) {
	session, err := r.getSessionFromRedis(ctx, sessionToken)
	if err != nil {
		if !errors.Is(err, dto.ErrExpiredSession) {
			return Session{}, err
		}
		err = r.db.WithContext(ctx).Where("token = ?", sessionToken).First(&session).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return Session{}, dto.ErrSessionNotFound
			}
			return Session{}, err
		}
	}
	return session, nil
}

// RefreshToken changes the access session.Token for a Session and increases lifespan of session
//
// Rewrites Session.PrevToken to oldToken for possible security check
// Rewrites Session.ExpiresAt to Now + duration
//
// returns dto.ErrSessionTokenConflict if newToken is used by another session
// returns gorm or redis error on internal errors
func (r *repository) RefreshToken(ctx context.Context, ses Session, newToken Token, duration time.Duration) (Session, error) {
	err := r.rdb.Del(ctx, ses.Token.GetRedisKey()).Err()
	if err != nil {
		return Session{}, err
	}
	ses.PrevToken = ses.Token
	ses.Token = newToken
	ses.ExpiresAt = time.Now().Add(duration)
	err = r.db.WithContext(ctx).Updates(&ses).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return Session{}, dto.ErrSessionTokenConflict
		}
		return Session{}, err
	}
	err = r.setSessionOnRedis(ctx, ses)
	if err != nil {
		return Session{}, err
	}
	return ses, nil
}

// Delete takes session.Token and closes that session permanently
//
// returns redis or gorm error on internal error
func (r *repository) Delete(ctx context.Context, token Token) error {
	err := r.db.WithContext(ctx).Where("token = ?", token).Delete(&Session{}).Error
	if err != nil {
		return err
	}
	err = r.rdb.Del(ctx, token.GetRedisKey()).Err()
	return err
}

// getSessionFromRedis takes session.Token for finding Session on redis
//
// returns dto.ErrExpiredSession if session is not in redis/expired.
// returns redis or json error on internal errors.
func (r *repository) getSessionFromRedis(ctx context.Context, token Token) (Session, error) {
	result, err := r.rdb.Get(ctx, token.GetRedisKey()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return Session{}, dto.ErrExpiredSession
		}
		return Session{}, err
	}
	var session Session
	err = json.Unmarshal([]byte(result), &session)
	return session, err
}

// setSessionOnRedis puts session to redis for fast access
//
// # It will add 15 minutes extra on expiration time for refresh
//
// returns redis error on internal errors
func (r *repository) setSessionOnRedis(ctx context.Context, session Session) error {
	return r.rdb.
		Set(ctx, session.Token.GetRedisKey(), session, time.Until(
			session.ExpiresAt.Add(15*time.Minute))).
		Err()
}

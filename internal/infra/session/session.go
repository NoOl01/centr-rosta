package session

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/pkg/logger"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RepositorySession interface {
	Create(ctx context.Context, session Session) (string, error)
	Get(ctx context.Context, sessionID string) (*Session, error)
	Update(ctx context.Context, sessionID string, session Session) error
	Delete(ctx context.Context, sessionID string) error
}

type sessionRepository struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewRepositorySession(client *redis.Client) RepositorySession {
	return &sessionRepository{
		rdb: client,
		ttl: 24 * time.Hour * 30,
	}
}

type Session struct {
	UserID       string
	DeviceToken  string
	AccessToken  string
	RefreshToken string
}

func (s *sessionRepository) Create(ctx context.Context, session Session) (string, error) {
	key := uuid.NewString()

	data, err := json.Marshal(session)
	if err != nil {
		logger.Log.Error(log_names.RedisSession, err.Error())
		return "", errs.InternalError
	}

	if err := s.rdb.Set(ctx, key, data, s.ttl).Err(); err != nil {
		logger.Log.Error(log_names.RedisSession, err.Error())
		return "", errs.InternalError
	}

	return key, nil
}

func (s *sessionRepository) Get(ctx context.Context, sessionID string) (*Session, error) {
	data, err := s.rdb.Get(ctx, sessionID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Log.Debug(log_names.RedisSession, errs.SessionNotFound.Error())
			return nil, errs.SessionNotFound
		}
		logger.Log.Error(log_names.RedisSession, err.Error())
		return nil, errs.InternalError
	}

	var session Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		logger.Log.Error(log_names.RedisSession, err.Error())
		return nil, errs.InternalError
	}

	return &session, nil
}

func (s *sessionRepository) Update(ctx context.Context, sessionID string, session Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		logger.Log.Error(log_names.RedisSession, err.Error())
		return errs.InternalError
	}

	if err := s.rdb.Set(ctx, sessionID, data, s.ttl).Err(); err != nil {
		logger.Log.Error(log_names.RedisSession, err.Error())
		return errs.InternalError
	}

	return nil
}

func (s *sessionRepository) Delete(ctx context.Context, sessionID string) error {
	if err := s.rdb.Del(ctx, sessionID).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Log.Debug(log_names.RedisSession, err.Error())
			return errs.SessionNotFound
		}
		logger.Log.Error(log_names.RedisSession, err.Error())
		return errs.InternalError
	}

	return nil
}

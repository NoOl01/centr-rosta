package redis

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/pkg/logger"
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *SessionRepository) Create(ctx context.Context, session entity.Session) (string, error) {
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

func (s *SessionRepository) Get(ctx context.Context, sessionID string) (*entity.Session, error) {
	data, err := s.rdb.Get(ctx, sessionID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Log.Debug(log_names.RedisSession, errs.SessionNotFound.Error())
			return nil, errs.SessionNotFound
		}
		logger.Log.Error(log_names.RedisSession, err.Error())
		return nil, errs.InternalError
	}

	var session entity.Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		logger.Log.Error(log_names.RedisSession, err.Error())
		return nil, errs.InternalError
	}

	return &session, nil
}

func (s *SessionRepository) Update(ctx context.Context, sessionID string, session entity.Session) error {
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

func (s *SessionRepository) Delete(ctx context.Context, sessionID string) error {
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

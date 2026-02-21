package session

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/pkg/logger"
	"context"
	"encoding/json"
	"errors"

	"github.com/redis/go-redis/v9"
)

func (s *sessionRepository) Get(ctx context.Context, sessionID string) (*Session, error) {
	data, err := s.rdb.Get(ctx, sessionID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Log.Error(log_names.RedisSession, errs.SessionNotFound.Error())
			return nil, errs.SessionNotFound
		}
		logger.Log.Error(log_names.RedisSession, err.Error())
		return nil, err
	}

	var session Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		logger.Log.Error(log_names.RedisSession, err.Error())
		return nil, err
	}

	return &session, nil
}

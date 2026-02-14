package session

import (
	"centr_rosta/internal/consts"
	"centr_rosta/pkg/logger"
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

func (s *sessionRepository) Create(ctx context.Context, session Session) (string, error) {
	key := uuid.NewString()

	data, err := json.Marshal(session)
	if err != nil {
		logger.Log.Error(consts.RedisSession, err.Error())
	}

	if err := s.rdb.Set(ctx, key, data, s.ttl).Err(); err != nil {
		return "", err
	}

	return key, nil
}

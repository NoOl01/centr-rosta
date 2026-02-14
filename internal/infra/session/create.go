package session

import (
	"centr_rosta/internal/consts"
	"centr_rosta/pkg/logger"
	"context"
	"encoding/json"
)

func (s *sessionRepository) Create(ctx context.Context, refreshToken string, session Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		logger.Log.Error(consts.RedisSession, err.Error())
	}

	if err := s.rdb.Set(ctx, refreshToken, data, s.ttl).Err(); err != nil {
		return err
	}

	return nil
}

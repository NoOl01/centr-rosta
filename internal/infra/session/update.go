package session

import (
	"centr_rosta/internal/consts"
	"centr_rosta/pkg/logger"
	"context"
	"encoding/json"
)

func (s *sessionRepository) Update(ctx context.Context, sessionID string, session Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		logger.Log.Error(consts.RedisSession, err.Error())
	}

	if err := s.rdb.Set(ctx, sessionID, data, s.ttl).Err(); err != nil {
		return err
	}

	return nil
}

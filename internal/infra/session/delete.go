package session

import (
	"centr_rosta/internal/consts"
	"centr_rosta/pkg/logger"
	"context"
)

func (s *sessionRepository) Delete(ctx context.Context, sessionID string) error {
	if err := s.rdb.Del(ctx, sessionID).Err(); err != nil {
		logger.Log.Error(consts.RedisSession, err.Error())
		return err
	}

	return nil
}

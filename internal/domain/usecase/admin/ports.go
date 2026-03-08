package admin

import (
	"centr_rosta/internal/domain/entity"
	"context"
	"time"
)

type ITransactionRepository interface {
	TransactionsByTimePeriod(from, to time.Time) ([]entity.Transaction, error)
}

type ISessionRepository interface {
	Get(ctx context.Context, sessionID string) (*entity.Session, error)
}

type IJwt interface {
	ValidateJwt(token string) (*entity.Payload, error)
}

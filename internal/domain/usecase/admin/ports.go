package admin

import (
	"centr_rosta/internal/domain/entity"
	"context"
	"time"
)

type TransactionRepository interface {
	TransactionsByTimePeriod(from, to time.Time) ([]entity.Transaction, error)
}

type SessionRepository interface {
	Get(ctx context.Context, sessionID string) (*entity.Session, error)
}

type Jwt interface {
	ValidateJwt(token string) (*entity.Payload, error)
}

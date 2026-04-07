package validate

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type SessionRepository interface {
	Get(ctx context.Context, sessionID string) (*entity.Session, error)
}

type Jwt interface {
	ValidateJwt(token string) (*entity.Payload, error)
}

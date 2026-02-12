package session

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RepositorySession interface {
	Create(ctx context.Context, session Session, ttl time.Duration) error
	Get(ctx context.Context, sessionID string) (*Session, error)
	Delete(ctx context.Context, sessionID string) error
}

type sessionRepository struct {
	client *redis.Client
}

func NewRepositorySession(client *redis.Client) RepositorySession {
	return &sessionRepository{
		client: client,
	}
}

type Session struct {
	UserID       int64
	RefreshToken string
	DeviceToken  string
}

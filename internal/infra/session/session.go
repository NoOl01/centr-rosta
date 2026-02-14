package session

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RepositorySession interface {
	Create(ctx context.Context, session Session) (string, error)
	Get(ctx context.Context, sessionID string) (*Session, error)
	Delete(ctx context.Context, sessionID string) error
}

type sessionRepository struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewRepositorySession(client *redis.Client) RepositorySession {
	return &sessionRepository{
		rdb: client,
		ttl: 24 * time.Hour * 30,
	}
}

type Session struct {
	UserID       string
	DeviceToken  string
	AccessToken  string
	RefreshToken string
}

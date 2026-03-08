package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewRepositorySession(client *redis.Client) *SessionRepository {
	return &SessionRepository{
		rdb: client,
		ttl: 24 * time.Hour * 30,
	}
}

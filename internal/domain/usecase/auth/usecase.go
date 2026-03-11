package auth

import (
	"centr_rosta/internal/domain/entity"
	"context"
	"strconv"
)

type UseCaseAuth interface {
	Register(ctx context.Context, user entity.User) (string, string, string, error)
	Login(ctx context.Context, user entity.Login) (string, string, string, error)
	Refresh(ctx context.Context, sessionID string, refreshData entity.Refresh) (string, string, error)
	Logout(ctx context.Context, sessionID string) error
	CheckAccess(ctx context.Context, sessionId, authToken string) error
}

type useCaseAuth struct {
	ur   UserRepository
	sr   SessionRepository
	jwt  Jwt
	pass PassHash
}

func NewUseCaseAuth(ur UserRepository, sr SessionRepository, jwt Jwt, pass PassHash) UseCaseAuth {
	return &useCaseAuth{
		ur:   ur,
		sr:   sr,
		jwt:  jwt,
		pass: pass,
	}
}

func (ua *useCaseAuth) createSession(ctx context.Context, userID int64, userRole string) (string, string, string, error) {
	userIDStr := strconv.FormatInt(userID, 10)

	newPayload := entity.Payload{
		UserId: userIDStr,
		Role:   userRole,
	}

	accessToken, refreshToken, err := ua.jwt.GenerateToken(newPayload)
	if err != nil {
		return "", "", "", err
	}

	newSession := entity.Session{
		UserID:       userIDStr,
		DeviceToken:  "",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	sessionId, err := ua.sr.Create(ctx, newSession)
	if err != nil {
		return "", "", "", err
	}

	return accessToken, refreshToken, sessionId, nil
}

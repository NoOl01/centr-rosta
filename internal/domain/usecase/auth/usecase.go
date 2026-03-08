package auth

import (
	"centr_rosta/internal/domain/entity"
	"context"
	"strconv"
)

type IUseCaseAuth interface {
	Register(ctx context.Context, user entity.User) (string, string, string, error)
	Login(ctx context.Context, user entity.Login) (string, string, string, error)
	Refresh(ctx context.Context, sessionID string, refreshData entity.Refresh) (string, string, error)
	Logout(ctx context.Context, sessionID string) error
	CheckAccess(ctx context.Context, sessionId, authToken string) error
}

type UseCaseAuth struct {
	ur   IUserRepository
	sr   ISessionRepository
	jwt  IJwt
	pass IPassHash
}

func NewUseCaseAuth(ur IUserRepository, sr ISessionRepository, jwt IJwt, pass IPassHash) IUseCaseAuth {
	return &UseCaseAuth{
		ur:   ur,
		sr:   sr,
		jwt:  jwt,
		pass: pass,
	}
}

func (ua *UseCaseAuth) createSession(ctx context.Context, userID int64, userRole string) (string, string, string, error) {
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

package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository/user"
	"centr_rosta/internal/utils/jwt"
	"context"
	"strconv"
)

type UseCaseAuth interface {
	Register(ctx context.Context, user dto.User) (string, string, string, error)
	Login(ctx context.Context, user dto.Login) (string, string, string, error)
	Refresh(ctx context.Context, sessionID string, refreshData dto.Refresh) (string, string, error)
	Logout(ctx context.Context, sessionID string) error
	CheckAccess(ctx context.Context, sessionId, authToken string) error
}

type useCaseAuth struct {
	ru      user.RepositoryUser
	session session.RepositorySession
}

func NewService(ru user.RepositoryUser, session session.RepositorySession) UseCaseAuth {
	return &useCaseAuth{
		ru:      ru,
		session: session,
	}
}

func (ua *useCaseAuth) createSession(ctx context.Context, userID int64, userRole string) (string, string, string, error) {
	userIDStr := strconv.FormatInt(userID, 10)

	newPayload := jwt.Payload{
		UserId: userIDStr,
		Role:   userRole,
	}

	accessToken, refreshToken, err := jwt.GenerateToken(newPayload)
	if err != nil {
		return "", "", "", err
	}

	newSession := session.Session{
		UserID:       userIDStr,
		DeviceToken:  "",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	sessionId, err := ua.session.Create(ctx, newSession)
	if err != nil {
		return "", "", "", err
	}

	return accessToken, refreshToken, sessionId, nil
}

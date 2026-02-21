package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/utils/jwt"
	"context"
)

func (ua *useCaseAuth) Refresh(ctx context.Context, sessionID string, refreshData dto.Refresh) (string, string, error) {
	payload, err := jwt.ValidateJwt(refreshData.RefreshToken)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := jwt.GenerateToken(*payload)
	if err != nil {
		return "", "", err
	}

	newSession := session.Session{
		UserID:       payload.UserId,
		DeviceToken:  "",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := ua.session.Update(ctx, sessionID, newSession); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

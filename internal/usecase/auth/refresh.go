package auth

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/utils/jwt"
	"centr_rosta/pkg/logger"
	"context"
)

func (ua *useCaseAuth) Refresh(ctx context.Context, sessionID string, refreshData dto.Refresh) (string, string, error) {
	logger.Log.Debug(log_names.UARefresh, "refreshing... get parameters: sessionID: "+sessionID+" refreshToken: "+refreshData.RefreshToken)
	logger.Log.Debug(log_names.UARefresh, "validating token")

	oldSession, err := ua.session.Get(ctx, sessionID)
	if err != nil {
		return "", "", err
	}

	if oldSession.RefreshToken != refreshData.RefreshToken {
		_ = ua.session.Delete(ctx, sessionID)
		return "", "", errs.InvalidToken
	}

	payload, err := jwt.ValidateJwt(refreshData.RefreshToken)
	if err != nil {
		_ = ua.session.Delete(ctx, sessionID)
		return "", "", err
	}

	logger.Log.Debug(log_names.UARefresh, "generating tokens")

	accessToken, refreshToken, err := jwt.GenerateToken(*payload)
	if err != nil {
		return "", "", err
	}

	logger.Log.Debug(log_names.UARefresh, "create new session")

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

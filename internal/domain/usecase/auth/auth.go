package auth

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/pkg/logger"
	"context"
	"strconv"
)

func (ua *useCaseAuth) Register(ctx context.Context, user entity.User) (string, string, string, error) {
	var err error
	*user.Password, err = ua.pass.EncryptPassword(*user.Password)
	if err != nil {
		return "", "", "", err
	}

	if err := ua.ur.CreateUser(&user); err != nil {
		return "", "", "", err
	}

	return ua.createSession(ctx, *user.ID, *user.Role)
}

func (ua *useCaseAuth) Login(ctx context.Context, user entity.Login) (string, string, string, error) {
	dbUser, err := ua.ur.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", "", err
	}

	if err := ua.pass.CheckPass(user.Password, *dbUser.Password); err != nil {
		return "", "", "", err
	}

	return ua.createSession(ctx, *dbUser.ID, *dbUser.Role)
}

func (ua *useCaseAuth) Refresh(ctx context.Context, sessionID string, refreshData entity.Refresh) (string, string, error) {
	logger.Log.Debug(log_names.UARefresh, "refreshing... get parameters: sessionID: "+sessionID+" refreshToken: "+refreshData.RefreshToken)
	logger.Log.Debug(log_names.UARefresh, "validating token")

	oldSession, err := ua.sr.Get(ctx, sessionID)
	if err != nil {
		return "", "", err
	}

	if oldSession.RefreshToken != refreshData.RefreshToken {
		_ = ua.sr.Delete(ctx, sessionID)
		return "", "", errs.InvalidToken
	}

	payload, err := ua.jwt.ValidateJwt(refreshData.RefreshToken)
	if err != nil {
		_ = ua.sr.Delete(ctx, sessionID)
		return "", "", err
	}

	logger.Log.Debug(log_names.UARefresh, "generating tokens")

	accessToken, refreshToken, err := ua.jwt.GenerateToken(*payload)
	if err != nil {
		return "", "", err
	}

	logger.Log.Debug(log_names.UARefresh, "create new redis")

	newSession := entity.Session{
		UserID:       payload.UserId,
		DeviceToken:  "",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := ua.sr.Update(ctx, sessionID, newSession); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (ua *useCaseAuth) CheckAccess(ctx context.Context, sessionId, authToken string) error {
	payload, err := ua.validate.Validate(ctx, sessionId, authToken)
	if err != nil {
		return err
	}

	userID, err := strconv.ParseInt(payload.UserId, 10, 64)
	if err != nil {
		return errs.InternalError
	}

	_, err = ua.ur.GetUserById(userID)
	if err != nil {
		return err
	}

	return nil
}

func (ua *useCaseAuth) Logout(ctx context.Context, sessionID string) error {
	return ua.sr.Delete(ctx, sessionID)
}

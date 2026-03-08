package auth

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/pkg/logger"
	"context"
	"strconv"
)

func (ua *UseCaseAuth) Register(ctx context.Context, user entity.User) (string, string, string, error) {
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

func (ua *UseCaseAuth) Login(ctx context.Context, user entity.Login) (string, string, string, error) {
	dbUser, err := ua.ur.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", "", err
	}

	if err := ua.pass.CheckPass(user.Password, *dbUser.Password); err != nil {
		return "", "", "", err
	}

	return ua.createSession(ctx, *dbUser.ID, *dbUser.Role)
}

func (ua *UseCaseAuth) Refresh(ctx context.Context, sessionID string, refreshData entity.Refresh) (string, string, error) {
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

func (ua *UseCaseAuth) CheckAccess(ctx context.Context, sessionId, authToken string) error {
	logger.Log.Debug(log_names.UACheckAccess, "checking access...")

	logger.Log.Debug(log_names.UACheckAccess, "getting redis from redis")
	sess, err := ua.sr.Get(ctx, sessionId)
	if err != nil {
		return err
	}
	if sess == nil {
		return errs.SessionNotFound
	}

	logger.Log.Debug(log_names.UACheckAccess, "comparing tokens")

	if authToken != sess.AccessToken {
		return errs.InvalidToken
	}

	logger.Log.Debug(log_names.UACheckAccess, "validating access token")

	payload, err := ua.jwt.ValidateJwt(authToken)
	if err != nil {
		logger.Log.Debug(log_names.UACheckAccess, "token is invalid. delete redis")
		return err
	}

	logger.Log.Debug(log_names.UACheckAccess, "parse userID")

	userID, err := strconv.ParseInt(payload.UserId, 10, 64)
	if err != nil {
		return errs.InternalError
	}

	logger.Log.Debug(log_names.UACheckAccess, "getting user from database")

	_, err = ua.ur.GetUserById(userID)
	if err != nil {
		logger.Log.Debug(log_names.UACheckAccess, "user not found. delete redis")
		return err
	}

	logger.Log.Debug(log_names.UACheckAccess, "check access passed successfully")
	return nil
}

func (ua *UseCaseAuth) Logout(ctx context.Context, sessionID string) error {
	return ua.sr.Delete(ctx, sessionID)
}

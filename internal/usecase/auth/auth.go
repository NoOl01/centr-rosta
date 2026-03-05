package auth

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository/models"
	"centr_rosta/internal/utils/jwt"
	"centr_rosta/internal/utils/pass_hash"
	"centr_rosta/pkg/logger"
	"context"
	"strconv"
)

func (ua *useCaseAuth) Register(ctx context.Context, user dto.User) (string, string, string, error) {
	var err error
	*user.Password, err = pass_hash.EncryptPassword(*user.Password)
	if err != nil {
		return "", "", "", err
	}

	newUser := models.User{
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
		Password:  *user.Password,
	}

	if err := ua.ru.CreateUser(&newUser); err != nil {
		return "", "", "", err
	}

	return ua.createSession(ctx, newUser.ID, newUser.Role)
}

func (ua *useCaseAuth) Login(ctx context.Context, user dto.Login) (string, string, string, error) {
	dbUser, err := ua.ru.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", "", err
	}

	if err := pass_hash.CheckPass(user.Password, dbUser.Password); err != nil {
		return "", "", "", err
	}

	return ua.createSession(ctx, dbUser.ID, dbUser.Role)
}

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

func (ua *useCaseAuth) CheckAccess(ctx context.Context, sessionId, authToken string) error {
	logger.Log.Debug(log_names.UACheckAccess, "checking access...")

	logger.Log.Debug(log_names.UACheckAccess, "getting session from redis")
	sess, err := ua.session.Get(ctx, sessionId)
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

	payload, err := jwt.ValidateJwt(authToken)
	if err != nil {
		logger.Log.Debug(log_names.UACheckAccess, "token is invalid. delete session")
		return err
	}

	logger.Log.Debug(log_names.UACheckAccess, "parse userID")

	userID, err := strconv.ParseInt(payload.UserId, 10, 64)
	if err != nil {
		return errs.InternalError
	}

	logger.Log.Debug(log_names.UACheckAccess, "getting user from database")

	_, err = ua.ru.GetUserById(userID)
	if err != nil {
		logger.Log.Debug(log_names.UACheckAccess, "user not found. delete session")
		return err
	}

	logger.Log.Debug(log_names.UACheckAccess, "check access passed successfully")
	return nil
}

func (ua *useCaseAuth) Logout(ctx context.Context, sessionID string) error {
	return ua.session.Delete(ctx, sessionID)
}

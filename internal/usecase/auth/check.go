package auth

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/utils/jwt"
	"centr_rosta/pkg/logger"
	"context"
	"strconv"
)

func (ua *useCaseAuth) CheckAccess(ctx context.Context, sessionId, authToken string) error {
	logger.Log.Debug(log_names.UACheckAccess, "checking access...")

	logger.Log.Debug(log_names.UACheckAccess, "getting session from redis")
	session, err := ua.session.Get(ctx, sessionId)
	if err != nil {
		return err
	}
	if session == nil {
		return errs.SessionNotFound
	}

	logger.Log.Debug(log_names.UACheckAccess, "comparing tokens")

	if authToken != session.AccessToken {
		_ = ua.deleteSession(ctx, sessionId)
		return errs.InvalidToken
	}

	logger.Log.Debug(log_names.UACheckAccess, "validating access token")

	payload, err := jwt.ValidateJwt(authToken)
	if err != nil {
		logger.Log.Debug(log_names.UACheckAccess, "token is invalid. delete session")
		_ = ua.session.Delete(ctx, sessionId)
		return err
	}

	logger.Log.Debug(log_names.UACheckAccess, "parse userID")

	userID, err := strconv.ParseInt(payload.UserId, 10, 64)
	if err != nil {
		return err
	}

	logger.Log.Debug(log_names.UACheckAccess, "getting user from database")

	_, err = ua.ru.GetUseById(userID)
	if err != nil {
		logger.Log.Debug(log_names.UACheckAccess, "user not found. delete session")
		_ = ua.deleteSession(ctx, sessionId)
		return err
	}

	logger.Log.Debug(log_names.UACheckAccess, "check access passed successfully")
	return nil
}

func (ua *useCaseAuth) deleteSession(ctx context.Context, sessionId string) error {
	return ua.session.Delete(ctx, sessionId)
}

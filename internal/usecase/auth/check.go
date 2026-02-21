package auth

import (
	"centr_rosta/internal/consts"
	"centr_rosta/internal/utils/jwt"
	"centr_rosta/pkg/logger"
	"context"
	"strconv"
)

func (ua *useCaseAuth) CheckAccess(ctx context.Context, sessionId, authToken string) error {
	logger.Log.Debug(consts.UACheckAccess, "checking access...")

	logger.Log.Debug(consts.UACheckAccess, "getting session from redis")
	session, err := ua.session.Get(ctx, sessionId)
	if err != nil {
		return err
	}
	if session == nil {
		return consts.SessionNotFound
	}

	logger.Log.Debug(consts.UACheckAccess, "comparing tokens")

	if authToken != session.AccessToken {
		_ = ua.deleteSession(ctx, sessionId)
		return consts.InvalidToken
	}

	logger.Log.Debug(consts.UACheckAccess, "validating access token")

	payload, err := jwt.ValidateJwt(authToken)
	if err != nil {
		logger.Log.Debug(consts.UACheckAccess, "token is invalid. delete session")
		_ = ua.session.Delete(ctx, sessionId)
		return err
	}

	logger.Log.Debug(consts.UACheckAccess, "parse userID")

	userID, err := strconv.ParseInt(payload.UserId, 10, 64)
	if err != nil {
		return err
	}

	logger.Log.Debug(consts.UACheckAccess, "getting user from database")

	_, err = ua.ru.GetUseById(userID)
	if err != nil {
		logger.Log.Debug(consts.UACheckAccess, "user not found. delete session")
		_ = ua.deleteSession(ctx, sessionId)
		return err
	}

	logger.Log.Debug(consts.UACheckAccess, "check access passed successfully")
	return nil
}

func (ua *useCaseAuth) deleteSession(ctx context.Context, sessionId string) error {
	return ua.session.Delete(ctx, sessionId)
}

package jwt

import (
	"centr_rosta/internal/config"
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	UserId string
	Role   string
}

func GenerateToken(payload Payload) (string, string, error) {
	key := []byte(config.Env.JwtSecret)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  payload.UserId,
		"role": payload.Role,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  payload.UserId,
		"role": payload.Role,
		"exp":  time.Now().Add(30 * 24 * time.Hour).Unix(),
	})

	accessTokenString, err := accessToken.SignedString(key)
	if err != nil {
		logger.Log.Error(log_names.JWT, err.Error())
		return "", "", errs.New(errs.InternalServerError, err)
	}

	refreshTokenString, err := refreshToken.SignedString(key)
	if err != nil {
		logger.Log.Error(log_names.JWT, err.Error())
		return "", "", errs.New(errs.InternalServerError, err)
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateJwt(token string) (*Payload, error) {
	jwtSecret := []byte(config.Env.JwtSecret)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Log.Debug(log_names.JWT, errs.UnexpectedSignMethod.Error())
			return nil, errs.UnexpectedSignMethod
		}
		return jwtSecret, nil
	})

	if err != nil {
		if parsedToken != nil && !parsedToken.Valid {
			logger.Log.Debug(log_names.JWT, errs.InvalidToken.Error())
			return nil, errs.InvalidToken
		}
		logger.Log.Error(log_names.JWT, err.Error())
		return nil, err
	}

	if parsedToken == nil || !parsedToken.Valid {
		logger.Log.Error(log_names.JWT, errs.InvalidToken.Error())
		return nil, errs.InvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		logger.Log.Error(log_names.JWT, errs.InvalidTokenClaims.Error())
		return nil, errs.InvalidTokenClaims
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		logger.Log.Error(log_names.JWT, errs.InvalidOrMissingClaim.Error())
		return nil, errs.InvalidOrMissingClaim
	}

	role, ok := claims["role"].(string)
	if !ok {
		logger.Log.Error(log_names.JWT, errs.InvalidOrMissingClaim.Error())
		return nil, errs.InvalidOrMissingClaim
	}

	return &Payload{
		UserId: userId,
		Role:   role,
	}, nil
}

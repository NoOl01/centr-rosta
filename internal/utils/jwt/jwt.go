package jwt

import (
	"centr_rosta/internal/config"
	"centr_rosta/internal/consts"
	"centr_rosta/pkg/logger"
	"errors"
	"fmt"
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
		logger.Log.Error(consts.JWT, err.Error())
		return "", "", err
	}

	refreshTokenString, err := refreshToken.SignedString(key)
	if err != nil {
		logger.Log.Error(consts.JWT, err.Error())
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateJwt(token string) (*Payload, error) {
	jwtSecret := []byte(config.Env.JwtSecret)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Log.Error(consts.JWT, consts.UnexpectedSignMethod.Error())
			return nil, consts.UnexpectedSignMethod
		}
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrInvalidKey) || !parsedToken.Valid {
			logger.Log.Error(consts.JWT, consts.InvalidToken.Error())
			return nil, consts.InvalidToken
		}
		logger.Log.Error(consts.JWT, err.Error())
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		logger.Log.Error(consts.JWT, consts.InvalidTokenClaimsType.Error())
		return nil, consts.InvalidTokenClaimsType
	}

	var payload Payload

	userId, ok := claims["sub"].(string)
	if !ok {
		logger.Log.Error(consts.JWT, consts.InvalidOrMissingClaim.Error())
		return nil, fmt.Errorf("%w: sub", consts.InvalidOrMissingClaim)
	}

	role, ok := claims["role"].(string)
	if !ok {
		logger.Log.Error(consts.JWT, consts.InvalidOrMissingClaim.Error())
		return nil, fmt.Errorf("%w: role", consts.InvalidOrMissingClaim)
	}

	payload.UserId = userId
	payload.Role = role

	return &payload, nil
}

func Refresh(refreshToken string) (string, string, error) {
	payload, err := ValidateJwt(refreshToken)
	if err != nil {
		logger.Log.Error(consts.JWT, err.Error())
		return "", "", err
	}

	return GenerateToken(*payload)
}

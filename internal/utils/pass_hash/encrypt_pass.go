package pass_hash

import (
	"centr_rosta/internal/consts"
	"centr_rosta/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error(consts.PassHash, err.Error())
		return "", err
	}

	return string(hash), nil
}

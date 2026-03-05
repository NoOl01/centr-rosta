package pass_hash

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/pkg/logger"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error(log_names.PassHash, err.Error())
		return "", errs.New(errs.InternalServerError, err)
	}

	return string(hash), nil
}

func CheckPass(password, dbPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			logger.Log.Debug(log_names.PassHash, errs.WrongPassword.Error())
			return errs.WrongPassword
		}
		logger.Log.Error(log_names.PassHash, err.Error())
		return errs.New(errs.InternalServerError, err)
	}
	return nil
}

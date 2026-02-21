package pass_hash

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/pkg/logger"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CheckPass(password, dbPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			logger.Log.Debug(log_names.PassHash, errs.WrongPassword.Error())
			return errs.WrongPassword
		}
		logger.Log.Error(log_names.PassHash, err.Error())
		return err
	}
	return nil
}

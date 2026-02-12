package pass_hash

import (
	"centr_rosta/internal/consts"
	"centr_rosta/pkg/logger"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CheckPass(password, dbPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			logger.Log.Debug(consts.PassHash, consts.WrongPassword.Error())
			return consts.WrongPassword
		}
		logger.Log.Error(consts.PassHash, err.Error())
		return err
	}
	return nil
}

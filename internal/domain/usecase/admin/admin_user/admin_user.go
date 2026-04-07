package admin_user

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/pkg/logger"
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

func (uau *useCaseAdminUser) GetUsers(ctx context.Context, sessionID, accessToken string) ([]entity.User, error) {
	if _, err := uau.validate.ValidateAdmin(ctx, sessionID, accessToken); err != nil {
		return nil, err
	}

	users, err := uau.ur.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uau *useCaseAdminUser) ResetPassword(ctx context.Context, sessionID, accessToken string, userID int64) (*string, error) {
	payload, err := uau.validate.ValidateAdmin(ctx, sessionID, accessToken)
	if err != nil {
		return nil, err
	}

	tokenUserID, err := strconv.ParseInt(payload.UserId, 10, 64)
	if err != nil {
		return nil, errs.InternalError
	}

	if tokenUserID == userID {
		return nil, errs.PermissionDenied
	}

	newPass, err := uau.generatePassword()
	if err != nil {
		return nil, err
	}

	hash, err := uau.pass.EncryptPassword(*newPass)
	if err != nil {
		return nil, err
	}

	if err := uau.ur.UpdateUser(userID, &entity.UpdateUser{
		Password: new(hash),
	}); err != nil {
		return nil, err
	}

	return newPass, nil
}

func (uau *useCaseAdminUser) UpdateRole(ctx context.Context, sessionID, accessToken, roleName string, userID int64) error {
	payload, err := uau.validate.ValidateAdmin(ctx, sessionID, accessToken)
	if err != nil {
		return err
	}

	tokenUserID, err := strconv.ParseInt(payload.UserId, 10, 64)
	if err != nil {
		return errs.InternalError
	}

	if tokenUserID == userID {
		return errs.PermissionDenied
	}

	return uau.ur.UpdateUserRole(userID, roleName)
}

var passAvailableSymbols = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"!", "@", "#", "$", "%", "^", "&", "*",
}

const passLength = 10

func (uau *useCaseAdminUser) generatePassword() (*string, error) {
	var sb strings.Builder
	symbolsLen := big.NewInt(int64(len(passAvailableSymbols)))

	for i := 0; i < passLength; i++ {
		idx, err := rand.Int(rand.Reader, symbolsLen)
		if err != nil {
			logger.Log.Error(log_names.PassGenerator, err.Error())
			return nil, errs.InternalError
		}
		sb.WriteString(passAvailableSymbols[idx.Int64()])
	}

	return new(sb.String()), nil
}

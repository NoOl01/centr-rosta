package admin

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/domain/entity"
	"context"
	"time"
)

func (uad *useCaseAdmin) TransactionStatsByTimePeriod(cxt context.Context, accessToken, sessionID, fromStr, toStr string) (*[]entity.Transaction, float64, error) {
	var from, to time.Time
	var err error

	session, err := uad.session.Get(cxt, sessionID)
	if err != nil {
		return nil, 0, err
	}
	if session == nil {
		return nil, 0, errs.SessionNotFound
	}

	if session.AccessToken != accessToken {
		return nil, 0, errs.InvalidToken
	}

	if _, err := uad.jwt.ValidateJwt(accessToken); err != nil {
		return nil, 0, err
	}

	if fromStr == "" {
		from = time.Now().AddDate(0, -1, 0)
	} else {
		from, err = time.Parse(timeLayout, fromStr)
		if err != nil {
			return nil, 0, errs.WrongTimeFormat
		}
	}

	if toStr == "" {
		to = time.Now()
	} else {
		to, err = time.Parse(timeLayout, toStr)
		if err != nil {
			return nil, 0, errs.WrongTimeFormat
		}
	}

	dbTransactions, err := uad.rt.TransactionsByTimePeriod(from, to)
	if err != nil {
		return nil, 0, err
	}

	var totalAmount float64

	transactions := make([]entity.Transaction, 0, len(dbTransactions))

	for _, tr := range dbTransactions {
		transactions = append(transactions, entity.Transaction{
			UserID: tr.UserID,
			User: entity.User{
				FirstName: tr.User.FirstName,
				LastName:  tr.User.LastName,
				Email:     tr.User.Email,
			},
			Amount:   tr.Amount,
			Type:     tr.Type,
			LessonID: tr.LessonID,
			Lesson: entity.Lesson{
				Name: tr.Lesson.Name,
			},
		})

		totalAmount += tr.Amount
	}

	return &transactions, totalAmount, nil
}

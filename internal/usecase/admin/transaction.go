package admin

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/dto"
	"context"
	"time"
)

func (uad *useCaseAdmin) TransactionStatsByTimePeriod(cxt context.Context, accessToken, sessionID, fromStr, toStr string) (*dto.TransactionStats, error) {
	var from, to time.Time
	var err error

	session, err := uad.session.Get(cxt, sessionID)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errs.SessionNotFound
	}

	if session.AccessToken != accessToken {
		return nil, errs.InvalidToken
	}

	if fromStr == "" {
		from = time.Now().AddDate(0, -1, 0)
	} else {
		from, err = time.Parse(timeLayout, fromStr)
		if err != nil {
			return nil, errs.WrongTimeFormat
		}
	}

	if toStr == "" {
		to = time.Now()
	} else {
		to, err = time.Parse(timeLayout, toStr)
		if err != nil {
			return nil, errs.WrongTimeFormat
		}
	}

	dbTransactions, err := uad.rt.TransactionsByTimePeriod(from, to)
	if err != nil {
		return nil, err
	}

	var totalAmount float64

	transactions := make([]dto.Transaction, 0, len(dbTransactions))

	for _, tr := range dbTransactions {
		transactions = append(transactions, dto.Transaction{
			UserID: tr.UserID,
			User: dto.UserInfo{
				FirstName: tr.User.FirstName,
				LastName:  tr.User.LastName,
				Email:     tr.User.Email,
			},
			Amount:   tr.Amount,
			Type:     tr.Type,
			LessonID: tr.LessonID,
			Lesson: dto.Lesson{
				Name: tr.Lesson.Name,
			},
		})

		totalAmount += tr.Amount
	}

	transactionStat := dto.TransactionStats{
		TotalAmount: totalAmount,
		Transaction: transactions,
	}

	return &transactionStat, nil
}

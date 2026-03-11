package admin

import (
	"centr_rosta/internal/consts/keys"
	dto2 "centr_rosta/internal/handler/dto"
	"centr_rosta/internal/handler/helper"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ha *HandlerAdmin) GetStatsByTimePeriod(c *gin.Context) {
	from := c.Query(keys.From)
	to := c.Query(keys.To)

	auth, _ := c.Get(keys.Authorization)
	sessionId, _ := c.Get(keys.XSessionID)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	uTransactions, totalAmount, err := ha.uad.TransactionStatsByTimePeriod(ctx, auth.(string), sessionId.(string), from, to)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	transactions := make([]dto2.Transaction, 0, len(*uTransactions))
	for _, tr := range *uTransactions {
		transactions = append(transactions, dto2.Transaction{
			UserID: tr.UserID,
			User: dto2.UserInfo{
				FirstName: tr.User.FirstName,
				LastName:  tr.User.LastName,
				Email:     tr.User.Email,
			},
			Amount:   tr.Amount,
			Type:     tr.Type,
			LessonID: tr.LessonID,
			Lesson: dto2.Lesson{
				Name: tr.Lesson.Name,
			},
		})
	}

	transactionsStat := dto2.TransactionStats{
		TotalAmount: totalAmount,
		Transaction: transactions,
	}

	c.JSON(http.StatusOK, dto2.Result{
		Result: transactionsStat,
		Error:  nil,
	})
}

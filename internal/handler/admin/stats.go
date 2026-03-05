package admin

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/pkg/logger"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ha *handlerAdmin) GetStatsByTimePeriod(c *gin.Context) {
	from := c.Query(keys.From)
	to := c.Query(keys.To)

	auth, _ := c.Get(keys.Authorization)
	sessionId, _ := c.Get(keys.XSessionID)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	transactionsStat, err := ha.uad.TransactionStatsByTimePeriod(ctx, auth.(string), sessionId.(string), from, to)
	if err != nil {
		var code int
		logger.Log.Debug(log_names.HAdStat, err.Error())
		if errors.Is(err, errs.WrongTimeFormat) {
			code = http.StatusBadRequest
		} else if errors.Is(err, errs.InvalidToken) {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}
		c.JSON(code, dto.Result{
			Result: nil,
			Error:  dto.Strconv(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: *transactionsStat,
		Error:  nil,
	})
}

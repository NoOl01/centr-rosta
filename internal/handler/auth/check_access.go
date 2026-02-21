package auth

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ha *handlerAuth) CheckAccess(c *gin.Context) {
	auth := c.GetHeader(keys.Authorization)
	sessionId := c.Query(keys.SessionId)

	if auth == "" {
		logger.Log.Debug(log_names.AuthHandler, errs.MissingHeader.Error())
		c.JSON(http.StatusUnauthorized, dto.Result{
			Error: dto.Strconv(errs.MissingHeader.Error()),
		})
		return
	}
	if sessionId == "" {
		logger.Log.Debug(log_names.AuthHandler, errs.MissingQueryParameter.Error())
		c.JSON(http.StatusUnauthorized, dto.Result{
			Error: dto.Strconv(errs.MissingQueryParameter.Error()),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	err := ha.ua.CheckAccess(ctx, sessionId, auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Result{
			Error: dto.Strconv(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: "ok",
	})
}

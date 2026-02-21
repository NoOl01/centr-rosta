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

func (ha *handlerAuth) Refresh(c *gin.Context) {
	sessionId := c.Query(keys.SessionId)
	if sessionId == "" {
		logger.Log.Debug(log_names.AuthHandler, errs.MissingQueryParameter.Error())
		c.JSON(http.StatusBadRequest, dto.Result{
			Error: dto.Strconv(errs.MissingQueryParameter.Error()),
		})
		return
	}

	var body dto.Refresh
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Log.Debug(log_names.AuthHandler, err.Error())
		c.JSON(http.StatusBadRequest, dto.Result{
			Error: dto.Strconv(errs.MissingQueryParameter.Error()),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 15*time.Second)
	defer cancel()

	accessToken, refreshToken, err := ha.ua.Refresh(ctx, sessionId, body)
	if err != nil {
		logger.Log.Debug(log_names.AuthHandler, err.Error())
		c.JSON(http.StatusInternalServerError, dto.Result{
			Error: dto.Strconv(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

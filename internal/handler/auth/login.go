package auth

import (
	"centr_rosta/internal/consts"
	"centr_rosta/internal/dto"
	"centr_rosta/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ha *handlerAuth) Login(c *gin.Context) {
	var body dto.Login
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Log.Debug(consts.AuthHandler, err.Error())
		c.JSON(http.StatusBadRequest, dto.Result{
			Error: dto.Strconv(err.Error()),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	accessToken, refreshToken, sessionId, err := ha.service.Login(ctx, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Result{
			Error: dto.Strconv(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"session_id":    sessionId,
		},
	})
}

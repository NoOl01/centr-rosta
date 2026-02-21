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

func (ha *handlerAuth) Logout(c *gin.Context) {
	sessionId := c.Query("session_id")
	if sessionId == "" {
		logger.Log.Debug(consts.AuthHandler, consts.MissingQueryParameter.Error())
		c.JSON(http.StatusBadRequest, dto.Result{
			Error: dto.Strconv(consts.MissingQueryParameter.Error()),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := ha.ua.Logout(ctx, sessionId); err != nil {
		logger.Log.Debug(consts.AuthHandler, err.Error())
		c.JSON(http.StatusInternalServerError, dto.Result{
			Error: dto.Strconv(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Error: nil,
	})
}

package auth

import (
	"centr_rosta/internal/consts"
	"centr_rosta/internal/consts/keys"
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
		logger.Log.Debug(consts.AuthHandler, consts.MissingHeader.Error())
		c.JSON(http.StatusUnauthorized, dto.Result{
			Error: dto.Strconv(consts.MissingHeader.Error()),
		})
		return
	}
	if sessionId == "" {
		logger.Log.Debug(consts.AuthHandler, consts.MissingQueryParameter.Error())
		c.JSON(http.StatusUnauthorized, dto.Result{
			Error: dto.Strconv(consts.MissingQueryParameter.Error()),
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

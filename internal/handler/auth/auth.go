package auth

import (
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ha *handlerAuth) Register(c *gin.Context) {
	handleAuth[dto.User](c, ha.ua.Register)
}

func (ha *handlerAuth) Login(c *gin.Context) {
	handleAuth[dto.Login](c, ha.ua.Login)
}

func handleAuth[T dto.User | dto.Login](c *gin.Context, uaFunc func(context.Context, T) (string, string, string, error)) {
	var body T
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	access, refresh, sid, err := uaFunc(ctx, body)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
		"session_id":    sid,
	})
}

func (ha *handlerAuth) Refresh(c *gin.Context) {
	logger.Log.Debug(log_names.HARefresh, "invoked refresh")

	sessionID, _ := c.Get(keys.SessionId)
	sessionIDVal, err := getHeaderVal(sessionID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Log.Debug(log_names.AuthHandler, "sessionID: "+sessionIDVal)

	var body dto.Refresh
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	logger.Log.Debug(log_names.HARefresh, "body: RefreshToken: "+body.RefreshToken)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 15*time.Second)
	defer cancel()

	accessToken, refreshToken, err := ha.ua.Refresh(ctx, sessionIDVal, body)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Log.Debug(log_names.HARefresh, "accessToken: "+accessToken+", refreshToken: "+refreshToken)

	c.JSON(http.StatusOK, dto.Result{
		Result: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

func (ha *handlerAuth) Logout(c *gin.Context) {
	sessionID, _ := c.Get(keys.SessionId)
	sessionIDVal, err := getHeaderVal(sessionID)
	if err != nil {
		handleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := ha.ua.Logout(ctx, sessionIDVal); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Error: nil,
	})
}

func (ha *handlerAuth) CheckAccess(c *gin.Context) {
	auth, _ := c.Get(keys.Authorization)
	sessionID, _ := c.Get(keys.SessionId)

	authVal, err := getHeaderVal(auth)
	if err != nil {
		handleError(c, err)
		return
	}

	sessionIDVal, err := getHeaderVal(sessionID)
	if err != nil {
		handleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	err = ha.ua.CheckAccess(ctx, sessionIDVal, authVal)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: "ok",
	})
}

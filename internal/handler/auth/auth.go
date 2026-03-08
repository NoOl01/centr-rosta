package auth

import (
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/entity"
	dto2 "centr_rosta/internal/handler/dto"
	"centr_rosta/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ha *HandlerAuth) Register(c *gin.Context) {
	var body dto2.User
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	dBody := entity.User{
		FirstName: *body.FirstName,
		LastName:  *body.LastName,
		Email:     *body.Email,
		Password:  body.Password,
	}

	access, refresh, sid, err := ha.ua.Register(ctx, dBody)
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

func (ha *HandlerAuth) Login(c *gin.Context) {
	var body dto2.Login
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	dBody := entity.Login{
		Email:    body.Email,
		Password: body.Password,
	}

	access, refresh, sid, err := ha.ua.Login(ctx, dBody)
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

func (ha *HandlerAuth) Refresh(c *gin.Context) {
	logger.Log.Debug(log_names.HARefresh, "invoked refresh")

	sessionID, _ := c.Get(keys.XSessionID)
	sessionIDVal, err := getHeaderVal(sessionID)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Log.Debug(log_names.AuthHandler, "sessionID: "+sessionIDVal)

	var body dto2.Refresh
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	logger.Log.Debug(log_names.HARefresh, "body: RefreshToken: "+body.RefreshToken)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 15*time.Second)
	defer cancel()

	dBody := entity.Refresh{
		RefreshToken: body.RefreshToken,
	}

	accessToken, refreshToken, err := ha.ua.Refresh(ctx, sessionIDVal, dBody)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Log.Debug(log_names.HARefresh, "accessToken: "+accessToken+", refreshToken: "+refreshToken)

	c.JSON(http.StatusOK, dto2.Result{
		Result: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

func (ha *HandlerAuth) Logout(c *gin.Context) {
	sessionID, _ := c.Get(keys.XSessionID)
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

	c.JSON(http.StatusOK, dto2.Result{
		Error: nil,
	})
}

func (ha *HandlerAuth) CheckAccess(c *gin.Context) {
	auth, _ := c.Get(keys.Authorization)
	sessionID, _ := c.Get(keys.XSessionID)

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

	c.JSON(http.StatusOK, dto2.Result{
		Result: "ok",
	})
}

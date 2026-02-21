package auth

import (
	"centr_rosta/internal/dto"
	"context"
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
		c.JSON(400, dto.Result{
			Error: dto.Strconv(err.Error()),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	access, refresh, sid, err := uaFunc(ctx, body)
	if err != nil {
		c.JSON(400, dto.Result{
			Error: dto.Strconv(err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
		"session_id":    sid,
	})
}

package auth

import (
	"centr_rosta/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

type HandlerAuth interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Refresh(c *gin.Context)
	Logout(c *gin.Context)
	CheckAccess(c *gin.Context)
}

type handlerAuth struct {
	ua auth.UseCaseAuth
}

func NewHandlerAuth(ua auth.UseCaseAuth) HandlerAuth {
	return &handlerAuth{ua: ua}
}

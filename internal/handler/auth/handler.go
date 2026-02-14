package auth

import (
	"centr_rosta/internal/service/auth"

	"github.com/gin-gonic/gin"
)

type HandlerAuth interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Refresh(c *gin.Context)
	Logout(c *gin.Context)
}

type handlerAuth struct {
	service auth.ServiceAuth
}

func NewHandlerAuth(serviceAuth auth.ServiceAuth) HandlerAuth {
	return &handlerAuth{service: serviceAuth}
}

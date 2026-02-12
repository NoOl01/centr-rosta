package auth

import (
	"centr_rosta/internal/service/auth"

	"github.com/gin-gonic/gin"
)

type HandlerAuth interface {
	Login(ctx *gin.Context)
}

type handlerAuth struct {
	service auth.ServiceAuth
}

func NewHandlerAuth(serviceAuth auth.ServiceAuth) HandlerAuth {
	return &handlerAuth{service: serviceAuth}
}

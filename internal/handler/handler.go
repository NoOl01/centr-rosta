package handler

import (
	"centr_rosta/internal/handler/auth"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Router(r *gin.Engine)
}

type handler struct {
	handlerAuth auth.HandlerAuth
}

func NewHandler(handlerAuth auth.HandlerAuth) Handler {
	return &handler{handlerAuth: handlerAuth}
}

package handler

import (
	"centr_rosta/internal/handler/admin"
	"centr_rosta/internal/handler/auth"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Router(r *gin.Engine)
}

type handler struct {
	handlerAuth  auth.HandlerAuth
	handlerAdmin admin.HandlerAdmin
}

func NewHandler(handlerAuth auth.HandlerAuth, handlerAdmin admin.HandlerAdmin) Handler {
	return &handler{
		handlerAuth:  handlerAuth,
		handlerAdmin: handlerAdmin,
	}
}
